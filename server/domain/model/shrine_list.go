package model

// 神社一覧画面のリクエストデータ定義
type ShrineListReq struct {
	Kinds       string `json:"kinds"`
	StdAreaCode string `json:"stdAreaCode"`
	Tag         string `json:"tag"`
}

// 神社一覧画面のレスポンスデータ定義
type ShrineListResp struct {
	Name            string   `json:"name"`
	Address         string   `json:"address"`
	PlusCode        string   `json:"plusCode"`
	PlaceID         string   `json:"placeId"`
	ObjectOfWorship []string `json:"objectOfWorship"`
	HasGoshuin      bool     `json:"hasGoshuin"`
}
