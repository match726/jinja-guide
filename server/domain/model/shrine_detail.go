package model

// 神社詳細画面のデータ定義
type ShrineDetails struct {
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
