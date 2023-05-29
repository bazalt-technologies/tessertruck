package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tracflow/pkg/models"
)

func (api *API) RulesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vars := mux.Vars(r)
		idString := vars["tractor_id"]
		id, err := strconv.Atoi(idString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		rules, err := api.s.GetRulesForTractor(id)
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
	case http.MethodPost:
		var rule models.Rule
		err := json.NewDecoder(r.Body).Decode(&rule)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = api.s.NewRule(rule)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodPatch:
		var rule models.Rule
		err := json.NewDecoder(r.Body).Decode(&rule)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = api.s.NewRule(rule)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		type request struct {
			ID int
		}
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = api.s.DeleteRule(req.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
