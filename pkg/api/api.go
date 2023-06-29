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
	api.r.HandleFunc("/api/v1/rules/{tractorID}", api.RulesHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions)
	api.r.HandleFunc("/api/v1/rules", api.RulesHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions)
	api.r.HandleFunc("/api/v1/notes", api.NotesHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions)
	api.r.HandleFunc("/api/v1/tractors", api.TractorsHandler).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

	api.r.Use(api.HeadersMiddleware)
}

func (api *API) Listen(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
