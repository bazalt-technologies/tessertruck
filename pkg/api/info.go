package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
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

	infoWriter(conn, id)
}

func infoWriter(conn *websocket.Conn, id int) {
	for {
		info := models.Randomize(id)
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
