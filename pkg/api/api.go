package api

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"tracflow/pkg/storage"
)

type API struct {
	r        *mux.Router
	s        *storage.Store
	upgrader websocket.Upgrader
}

func New(r *mux.Router, u websocket.Upgrader, store *storage.Store) *API {
	return &API{r: r, upgrader: u, s: store}
}

func (api *API) Handle() {
	api.r.HandleFunc("/api/v1/info/{id}", api.InfoHandler)
	api.r.HandleFunc("/api/v1/rules/{tractor_id}", api.RulesHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete)
	api.r.HandleFunc("/api/v1/rules", api.RulesHandler).Methods(http.MethodPost, http.MethodPatch, http.MethodDelete)
	api.r.HandleFunc("/api/v1/notes/{tractor_id}", api.NotesHandler).Methods(http.MethodGet)
	api.r.HandleFunc("/api/v1/notes", api.NotesHandler).Methods(http.MethodPost, http.MethodPatch, http.MethodDelete)
	api.r.HandleFunc("/api/v1/tractors", api.TractorsHandler).Methods(http.MethodGet, http.MethodPost)
}

func (api *API) Listen(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
