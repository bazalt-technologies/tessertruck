package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"time"
	"tracflow/pkg/models"
)

func (api *API) InfoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := api.upgrader.Upgrade(w, r, nil)
	log.Println(conn)
	if err != nil {
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, err.Error()))
		return
	}

	defer conn.Close()
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	rules, _ := api.s.GetRulesForTractor(id)
	infoWriter(conn, rules, id, api)
}

func infoWriter(conn *websocket.Conn, rules []models.Rule, id int, api *API) {
	for {
		var wg sync.WaitGroup
		info := models.Randomize(id)
		wg.Add(1)
		info.Teledata.RulesWarnings = checkRulesWarnings(info, rules, api)

		data, err := json.Marshal(info)
		if err != nil {
			log.Println(err.Error())
		}

		if err := conn.WriteMessage(1, data); err != nil {
			return
		}

		time.Sleep(time.Second * 3)
	}
}

func checkRulesWarnings(info models.Tractor, rules []models.Rule, api *API) []string {
	var ruleWarns []string
	for _, rule := range rules {
		if reflect.ValueOf(&info.Teledata).Elem().FieldByName(rule.FieldName).CanInt() {
			if int(reflect.ValueOf(&info.Teledata).Elem().FieldByName(rule.FieldName).Int()) > rule.ValInt {
				ruleWarns = append(ruleWarns, rule.FieldName)
				go writeNote(rule,
					info,
					int(reflect.ValueOf(&info.Teledata).Elem().FieldByName(rule.FieldName).Int()),
					api)
			}
		}
		if reflect.ValueOf(&info.Teledata).Elem().FieldByName(rule.FieldName).CanFloat() {
			if reflect.ValueOf(&info.Teledata).Elem().FieldByName(rule.FieldName).Float() > rule.ValFloat {
				ruleWarns = append(ruleWarns, rule.FieldName)
				go writeNote(rule,
					info,
					reflect.ValueOf(&info.Teledata).Elem().FieldByName(rule.FieldName).Float(),
					api)
			}
		}
	}
	return ruleWarns
}

type Num interface {
	int | float64
}

func writeNote[T Num](rule models.Rule, t models.Tractor, val T, api *API) {
	var field string
	switch rule.FieldName {
	case "SpeedRT":
		field = "скорости"
	case "EngineRPS":
		field = "оборотов в секунду"
	case "FuelLevel":
		field = "уровня топлива"
	case "FuelSpent":
		field = "количества израсходованного топлива"
	case "FuelConsumption":
		field = "потребления топлива"
	case "OilTemperature":
		field = "температуры масла"
	case "EnvTemperature":
		field = "температуры окр. среды"
	}

	var mess string
	switch val.(type) {
	case int:
		mess = fmt.Sprintf("Нарушение для трактора #%s: значение %s превысил(-о) допустимые значения:%d вместо %d",
			t.Name, field, val, rule.ValInt)
	case float64:
		mess = fmt.Sprintf("Нарушение для трактора #%s: значение %s превысил(-о) допустимые значения:%f вместо %f",
			t.Name, field, val, rule.ValFloat)
	}
	note := models.Note{Time: time.Now().Format("02-01-2006 15:04:05 (Мск)"), Message: mess}
	api.s.NewNote(note)
}
