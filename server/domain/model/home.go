package model

// ホーム画面のレスポンスデータ定義
type HomeContentsResp struct {
	RandomShrines struct {
		Name            string   `json:"name"`
		Furigana        string   `json:"furigana"`
		Address         string   `json:"address"`
		PlusCode        string   `json:"plusCode"`
		PlaceId         string   `json:"placeId"`
		ObjectOfWorship []string `json:"objectOfWorship"`
		Description     string   `json:"description"`
	} `json:"randomShrines"`
}
