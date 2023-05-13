package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"tracflow/pkg/api"
)

func main() {
	a := api.New(mux.NewRouter(), websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	})
	a.Handle()
	a.Listen("server:8082")
}
