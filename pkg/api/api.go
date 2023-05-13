package api

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

type API struct {
	r        *mux.Router
	upgrader websocket.Upgrader
}

func New(r *mux.Router, u websocket.Upgrader) *API {
	return &API{r: r, upgrader: u}
}

func (api *API) Handle() {
	api.r.HandleFunc("/api/v1/info/{id}", api.InfoHandler)
}

func (api *API) Listen(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
