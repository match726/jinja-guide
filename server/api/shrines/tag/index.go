package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/match726/jinja-guide/tree/main/server/models"
	"github.com/match726/jinja-guide/tree/main/server/trace"
)

func ShrinesTagHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		FetchShrineTagList(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func FetchShrineTagList(w http.ResponseWriter, r *http.Request) {

	var pg *models.Postgres
	var err error

	fmt.Println("FetchShrineTagList")
	// Contextを生成
	ctx := r.Context()
	shutdown, err := trace.InitTracerProvider()
	if err != nil {
		panic(err)
	}
	defer shutdown(ctx)
	ctx = trace.GetContextWithTraceID(r.Context(), "FetchShrineTagList")

	// HTTPリクエストからカスタムヘッダーを取得
	strCustom := r.Header.Get("ShrGuide-Shrines-Authorization")

	var tag string
	err = json.Unmarshal([]byte(strCustom), &tag)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	tagDecoded, _ := url.QueryUnescape(tag)
	fmt.Println(tagDecoded)

	pg, err = models.NewPool(ctx)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer pg.ClosePool(ctx)

	var shrlrs []*models.ShrinesListResp
	shrlrs, err = pg.GetShrinesListByTag(ctx, tagDecoded)
	if err != nil {
		fmt.Printf("[Err] <GetShrinesListByTag> Err:%s\n", err)
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
