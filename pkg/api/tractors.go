package api

import (
	"encoding/json"
	"net/http"
	"tracflow/pkg/models"
)

func (api *API) TractorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		tractors, err := api.s.GetTractors([]int{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(tractors)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(data)
	case http.MethodPost:
		var tractor models.Tractor
		err := json.NewDecoder(r.Body).Decode(&tractor)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = api.s.NewTractor(tractor)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//TODO
		//case http.MethodDelete:
		//
		//}
	}
}
