package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/match726/jinja-guide/tree/main/server/logger"
	"github.com/match726/jinja-guide/tree/main/server/models"
	"github.com/match726/jinja-guide/tree/main/server/trace"
)

type ShrinePostReq struct {
	Name         string `json:"name"`
	Furigana     string `json:"furigana"`
	Address      string `json:"address"`
	WikipediaURL string `json:"wikipediaUrl"`
}

func ShrineRegistHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodPost:
		RegisterShrine(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func RegisterShrine(w http.ResponseWriter, r *http.Request) {

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

	// HTTPリクエストからボディを取得
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	// ShrinePostReq構造体へ変換
	var shr models.Shrine
	var shrpr *ShrinePostReq
	err = json.Unmarshal([]byte(string(body)), &shrpr)
	if err != nil {
		logger.Error(ctx, "パラメータ取得失敗", "errmsg", err)
	}

	// Shrine構造体へ代入
	shr.ShrineName(shrpr.Name)
	shr.ShrineAddress(shrpr.Address)

	pg, err = models.NewPool(ctx)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer pg.ClosePool(ctx)

	// 住所より都道府県の取得
	prefname := models.ExtractPrefName(shr.Address)
	// 該当の都道府県の標準地域コード一覧を取得
	err = pg.GetStdAreaCodeByPrefName(ctx, prefname, &shr)
	if err != nil {
		logger.Error(ctx, "標準地域コード一覧取得失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// PlaceAPIから位置情報(PlaceID、緯度、経度)、及び取得した緯度経度からPlusCodeを取得
	err = models.GetLocnInfoFromPlaceAPI(ctx, &shr)
	if err != nil {
		logger.Error(ctx, "PlaceAPI取得失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// 取得した情報を元に神社登録
	err = pg.InsertShrine(ctx, &shr)
	if err != nil {
		logger.Error(ctx, "神社登録失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// 神社詳細情報を登録
	if len(shr.PlusCode) != 0 {
		if len(shrpr.Furigana) != 0 {
			err = pg.InsertShrineContents(ctx, 1, shrpr.Furigana, shr.PlusCode, 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[振り仮名]登録失敗", "errmsg", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		if len(shrpr.WikipediaURL) != 0 {
			err = pg.InsertShrineContents(ctx, 10, shrpr.WikipediaURL, shr.PlusCode, 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[WikipediaURL]登録失敗", "errmsg", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}

	writeJsonResp(w, &shr)

}

func writeJsonResp(w http.ResponseWriter, shr *models.Shrine) {

	type ShrinePostResp struct {
		Name          string `json:"name"`
		Address       string `json:"address"`
		StdAreaCode   string `json:"std_area_code"`
		PlusCode      string `json:"plus_code"`
		Seq           string `json:"seq"`
		PlaceID       string `json:"place_id"`
		GoogleMapLink string `json:"google_map_link"`
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	shrResp := &ShrinePostResp{shr.Name, shr.Address, shr.StdAreaCode, shr.PlusCode, shr.Seq, shr.PlaceID, "https://www.google.com/maps/search/?api=1&query=" + shr.Name + "&query_place_id=" + shr.PlaceID}
	b, err := json.Marshal(shrResp)
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
