package api

import (
	"encoding/json"
	"net/http"

	"github.com/match726/jinja-guide/tree/main/server/models"
)

func StdAreaCodeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		FetchStdAreaCodes(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func FetchStdAreaCodes(w http.ResponseWriter, r *http.Request) {

	sacs := models.GetAllStdAreaCodesFromEstat

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sacs)

}
