package model

// ホーム画面のレスポンスデータ定義
type HomeContentsResp struct {
	Shrines []*RandomShrines `json:"shrines"`
	Tags    []*AllTags       `json:"tags"`
}

// ホーム画面の神社表示用レスポンスデータ定義
type RandomShrines struct {
	Name            string   `json:"name"`
	Furigana        string   `json:"furigana"`
	Address         string   `json:"address"`
	PlusCode        string   `json:"plusCode"`
	PlaceId         string   `json:"placeId"`
	ObjectOfWorship []string `json:"objectOfWorship"`
	Description     string   `json:"description"`
}

// ホーム画面の関連ワード表示用レスポンスデータ定義
type AllTags struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
