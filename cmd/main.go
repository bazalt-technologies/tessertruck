package main

import (
	"fmt"
	"reflect"
	"tracflow/pkg/models"
)

func main() {
	//a := api.New(mux.NewRouter(), websocket.Upgrader{
	//	ReadBufferSize:  1024,
	//	WriteBufferSize: 1024,
	//	CheckOrigin: func(r *http.Request) bool {
	//		return true
	//	},
	//})
	//a.Handle()
	//a.Listen("server:8082")

	rule := models.Rule{FieldName: "SpeedRT", ValFloat: 60}
	for {
		s := models.Randomize(1)
		if reflect.ValueOf(&s).Elem().FieldByName(rule.FieldName).CanInt() {
			if int(reflect.ValueOf(&s).Elem().FieldByName(rule.FieldName).Int()) > rule.ValInt {
				fmt.Println(s)
			}
		}
		if reflect.ValueOf(&s).Elem().FieldByName(rule.FieldName).CanFloat() {
			if reflect.ValueOf(&s).Elem().FieldByName(rule.FieldName).Float() > rule.ValFloat {
				fmt.Println(s)

			}
		}
	}
}
