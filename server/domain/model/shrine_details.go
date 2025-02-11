package model

// 神社詳細画面のリクエストデータ定義
type ShrineDetailsReq struct {
	PlusCode string `json:"plusCode"`
}

// 神社詳細画面のレスポンスデータ定義
type ShrineDetailsResp struct {
	Name            string   `json:"name"`
	Furigana        string   `json:"furigana"`
	Image           string   `json:"image"`
	AltName         []string `json:"altName"`
	Address         string   `json:"address"`
	PlaceID         string   `json:"placeId"`
	Description     string   `json:"description"`
	Tags            []string `json:"tags"`
	FoundedYear     string   `json:"foundedYear"`
	ObjectOfWorship []string `json:"objectOfWorship"`
	ShrineRank      []string `json:"shrineRank"`
	HasGoshuin      bool     `json:"hasGoshuin"`
	WebsiteURL      string   `json:"websiteUrl"`
	WikipediaURL    string   `json:"wikipediaUrl"`
}
