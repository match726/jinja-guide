package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/match726/jinja-guide/tree/main/server/models"
)

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
	var sacs []models.StdAreaCode

	// HTTPリクエストからボディを取得
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	// Shrine構造体へ変換
	var shr models.Shrine
	err = json.Unmarshal([]byte(string(body)), &shr)
	if err != nil {
		fmt.Printf("[Err] RegisterShrine: パラメータ取得エラー, Err: %s\n", err)
	}

	pg, err = models.NewPool()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer pg.ClosePool()

	// 住所より都道府県の取得
	prefname := models.ExtractPrefName(shr.Address)
	// 該当の都道府県の標準地域コード一覧を取得
	sacs, err = pg.GetStdAreaCodeListByPrefName(prefname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// 標準地域コードの紐づけ
	for i := len(sacs) - 1; i >= 0; i-- {
		if sacs[i].MunicName1 == "" && sacs[i].MunicName2 == "" {
			continue
		} else {
			keyword := sacs[i].PrefName + sacs[i].MunicName1 + sacs[i].MunicName2
			if strings.HasPrefix(shr.Address, keyword) {
				shr.StdAreaCode = sacs[i].StdAreaCode
				break
			}
		}
	}

	// PlaceAPIから位置情報(PlaceID、緯度、経度)、取得した緯度、経度からPlusCodeを取得
	// ★画像を取得
	err = models.GetLocnInfoFromPlaceAPI(shr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = pg.InsertShrine(shr)
	if err != nil {
		fmt.Printf("[Err] InsertShrine: t_shrinesへのINSERT失敗, Err:%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}
