package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/match726/jinja-guide/tree/main/server/models"
)

func ShrinesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		FetchShrineDetails(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func FetchShrineDetails(w http.ResponseWriter, r *http.Request) {

	var pg *models.Postgres
	var err error

	// HTTPリクエストからカスタムヘッダーを取得
	strCustom := r.Header.Get("ShrGuide-Shrines-Authorization")

	var plus_code string
	err = json.Unmarshal([]byte(strCustom), &plus_code)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	pg, err = models.NewPool()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer pg.ClosePool()

	var shrd *models.ShrineDetails
	shrd, err = pg.GetShrineDetails(plus_code)
	if err != nil {
		fmt.Printf("[Err] <GetShrinesByStdAreaCode> Err:%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Println(shrd)
		writejsonResp(w, shrd)
	}

}

func writejsonResp(w http.ResponseWriter, shrd *models.ShrineDetails) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(shrd)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		s := `{"status":"500 Internal Server Error"}`
		if _, err := w.Write([]byte(s)); err != nil {
			fmt.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println(err)
	}

}
