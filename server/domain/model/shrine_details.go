package model

// 神社詳細画面のリクエストデータ定義
type ShrineDetailsReq struct {
	PlusCode string `json:"plusCode"`
}

// 神社詳細画面のレスポンスデータ定義
type ShrineDetailsResp struct {
	Name            string  `json:"name"`
	Furigana        Field   `json:"furigana"`
	Image           string  `json:"image"`
	AltName         []Field `json:"altName"`
	Address         string  `json:"address"`
	PlaceID         string  `json:"placeId"`
	Description     Field   `json:"description"`
	Tags            []Field `json:"tags"`
	FoundedYear     Field   `json:"foundedYear"`
	ObjectOfWorship []Field `json:"objectOfWorship"`
	ShrineRank      []Field `json:"shrineRank"`
	HasGoshuin      Field   `json:"hasGoshuin"`
	WebsiteURL      Field   `json:"websiteUrl"`
	WikipediaURL    Field   `json:"wikipediaUrl"`
}
