package model

// 神社一括登録テーブル定義
type ShrineRegister struct {
	Name             string     `json:"name"`
	Address          string     `json:"address"`
	Furigana         string     `json:"furigana"`
	AltNames         []string   `json:"altName"`
	Tags             []string   `json:"tag"`
	FoundedYear      []string   `json:"foundedYear"`
	ObjectOfWorships []string   `json:"objectOfWorship"`
	ShrineRanks      [][]string `json:"shrineRank"`
	HasGoshuin       string     `json:"hasGoshuin"`
	WebsiteURL       string     `json:"websiteUrl"`
	WikipediaURL     string     `json:"wikipediaUrl"`
}

// 神社／神社詳細情報登録画面のリクエストデータ定義
type ShrineRegisterReq struct {
	Name             []Field `json:"name"`
	Address          []Field `json:"address"`
	PlusCode         []Field `json:"plusCode"`
	PlaceID          []Field `json:"placeId"`
	Furigana         []Field `json:"furigana"`
	AltNames         []Field `json:"altName"`
	Tags             []Field `json:"tag"`
	FoundedYear      []Field `json:"foundedYear"`
	ObjectOfWorships []Field `json:"objectOfWorship"`
	ShrineRanks      []Field `json:"shrineRank"`
	HasGoshuin       []Field `json:"hasGoshuin"`
	WebsiteURL       []Field `json:"websiteUrl"`
	WikipediaURL     []Field `json:"wikipediaUrl"`
}

type Field struct {
	Id       int    `json:"id"`
	Seq      string `json:"seq"`
	Content1 string `json:"content1"`
	Content2 string `json:"content2"`
	Content3 string `json:"content3"`
}
