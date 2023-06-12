package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (api *API) NotesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(r.URL.Query().Get("tractorID"))
		if err != nil || id == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rules, err := api.s.GetNotesForTractor(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(rules)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}
