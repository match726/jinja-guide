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
	case http.MethodPost:
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
		fmt.Println("Message: データベース接続不可")
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer pg.Close()

	sacs, err = models.GetStdAreaCodes(pg)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sacs)

}

func UpdateStdAreaCodes(w http.ResponseWriter, r *http.Request) {

	var sacs models.StdAreaCodes
	sacs = models.GetAllStdAreaCodesFromEstat()
	fmt.Println(sacs)

}
