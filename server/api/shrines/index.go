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
		FetchShrineList(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func FetchShrineList(w http.ResponseWriter, r *http.Request) {

	var pg *models.Postgres
	var err error

	// HTTPリクエストからカスタムヘッダーを取得
	strCustom := r.Header.Get("ShrGuide-Shrines-Authorization")

	// SacRelationship構造体へ変換
	var sacr *models.SacRelationship
	err = json.Unmarshal([]byte(strCustom), &sacr)
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

	var shrs []*models.Shrine
	shrs, err = pg.GetShrinesByStdAreaCode(sacr)
	if err != nil {
		fmt.Printf("[Err] <GetShrinesByStdAreaCode> Err:%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Println(shrs)
		writejsonResp(w, shrs)
	}

}

func writejsonResp(w http.ResponseWriter, shrs []*models.Shrine) {

	type ShrinesListResp struct {
		Name            string   `json:"name"`
		Address         string   `json:"address"`
		PlusCode        string   `json:"plusCode"`
		PlaceID         string   `json:"placeId"`
		ObjectOfWorship []string `json:"objectOfWorship"`
		HasGoshuin      bool     `json:"hasGoshuin"`
	}

	var shrListResp []ShrinesListResp

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	for _, shr := range shrs {
		shrListResp = append(shrListResp, ShrinesListResp{
			Name:            shr.Name,
			Address:         shr.Address,
			PlusCode:        shr.PlusCode,
			PlaceID:         shr.PlaceID,
			ObjectOfWorship: nil,
			HasGoshuin:      false,
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
