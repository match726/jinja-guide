package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/match726/jinja-guide/tree/main/server/models"
	"github.com/match726/jinja-guide/tree/main/server/trace"
)

func ShrinesHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		FetchShrineList(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func FetchShrineList(w http.ResponseWriter, r *http.Request) {

	var pg *models.Postgres
	var err error

	// Contextを生成
	ctx := r.Context()
	shutdown, err := trace.InitTracerProvider()
	if err != nil {
		panic(err)
	}
	defer shutdown(ctx)
	ctx = trace.GetContextWithTraceID(r.Context(), "FetchShrineDetails")

	// HTTPリクエストからカスタムヘッダーを取得
	strCustom := r.Header.Get("ShrGuide-Shrines-Authorization")

	// SacRelationship構造体へ変換
	var sacr *models.SacRelationship
	err = json.Unmarshal([]byte(strCustom), &sacr)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	pg, err = models.NewPool(ctx)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer pg.ClosePool(ctx)

	var shrlrs []*models.ShrinesListResp
	shrlrs, err = pg.GetShrinesListByStdAreaCode(ctx, sacr)
	if err != nil {
		fmt.Printf("[Err] <GetShrinesListByStdAreaCode> Err:%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Println(shrlrs)
		writejsonResp(w, shrlrs)
	}

}

func writejsonResp(w http.ResponseWriter, shrlrs []*models.ShrinesListResp) {

	var shrListResp []models.ShrinesListResp

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	for _, shrlr := range shrlrs {
		shrListResp = append(shrListResp, models.ShrinesListResp{
			Name:            shrlr.Name,
			Address:         shrlr.Address,
			PlusCode:        shrlr.PlusCode,
			PlaceID:         shrlr.PlaceID,
			ObjectOfWorship: nil,
			HasGoshuin:      shrlr.HasGoshuin,
		})
	}

	b, err := json.Marshal(shrListResp)
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
