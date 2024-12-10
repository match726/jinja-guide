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

	var pg *models.Postgres
	var err error
	var sacs []models.StdAreaCodeGet

	// Contextを生成
	ctx := r.Context()

	pg, err = models.NewPool(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer pg.ClosePool(ctx)

	sacs, err = pg.GetStdAreaCodeList(ctx)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sacs)

}

func UpdateStdAreaCodes(w http.ResponseWriter, r *http.Request) {

	var pg *models.Postgres
	var err error

	// Contextを生成
	ctx := r.Context()

	pg, err = models.NewPool(ctx)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer pg.ClosePool(ctx)

	err = pg.UpdateStdAreaCode(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
