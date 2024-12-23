package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/match726/jinja-guide/tree/main/server/models"
	"github.com/match726/jinja-guide/tree/main/server/trace"
)

func PrefsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		FetchSacRelationship(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func FetchSacRelationship(w http.ResponseWriter, r *http.Request) {

	var pg *models.Postgres
	var err error

	// Contextを生成
	ctx := r.Context()
	err = trace.InitTracer()
	trace.ShowTraceID(ctx)
	pg, err = models.NewPool(ctx)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer pg.ClosePool(ctx)

	sacr, err := pg.GetSacRelationship(ctx)
	if err != nil {
		fmt.Printf("[Err] <GetSacRelationship> Err:%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		writejsonResp(w, sacr)
	}

}

func writejsonResp(w http.ResponseWriter, sacr []models.SacRelationship) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(sacr)
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
