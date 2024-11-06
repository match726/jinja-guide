package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/match726/jinja-guide/tree/main/server/models"
)

func StdAreaCodeHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		FetchStdAreaCodes(w, r)
	case http.MethodPut:
		UpdateStdAreaCodes(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func FetchStdAreaCodes(w http.ResponseWriter, r *http.Request) {

	var sacs models.StdAreaCodes

	pg, err := models.NewPool()
	if err != nil {
		fmt.Fprintln(w, fmt.Sprintf("<p>%s</p>", err))
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer pg.ClosePool()

	sacs, err = pg.GetStdAreaCodes()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sacs)

}

func UpdateStdAreaCodes(w http.ResponseWriter, r *http.Request) {

	var pg *models.Postgres

	pg, err := models.NewPool()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer pg.ClosePool()

	err = pg.UpdateStdAreaCode()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
