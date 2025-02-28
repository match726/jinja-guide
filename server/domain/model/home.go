package model

// ホーム画面のレスポンスデータ定義
type HomeContentsResp struct {
	Shrines []RandomShrines `json:"shrines"`
}

// ホーム画面の神社一覧表示用レスポンスデータ定義
type RandomShrines struct {
	Name            string   `json:"name"`
	Furigana        string   `json:"furigana"`
	Address         string   `json:"address"`
	PlusCode        string   `json:"plusCode"`
	PlaceId         string   `json:"placeId"`
	ObjectOfWorship []string `json:"objectOfWorship"`
	Description     string   `json:"description"`
}
