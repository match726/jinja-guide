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
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func FetchStdAreaCodes(w http.ResponseWriter, r *http.Request) {

	fmt.Println("FetchStdAreaCodes呼び出し前")
	var sacs models.StdAreaCodes
	sacs = models.GetAllStdAreaCodesFromEstat()
	fmt.Println("FetchStdAreaCodes呼び出し後")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sacs)

}
