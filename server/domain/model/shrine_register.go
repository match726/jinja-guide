package model

type Field struct {
	Id    int    `json:"id"`
	Seq   string `json:"seq"`
	Value string `json:"value"`
}

// 神社登録画面のリクエストデータ定義
// 神社一括登録テーブル定義
type ShrineRegisterReq struct {
	Name            string   `json:"name"`
	Address         string   `json:"address"`
	Furigana        string   `json:"furigana"`
	AltName         []string `json:"altName"`
	Tags            []string `json:"tags"`
	FoundedYear     string   `json:"foundedYear"`
	ObjectOfWorship []string `json:"objectOfWorship"`
	HasGoshuin      string   `json:"hasGoshuin"`
	WebsiteURL      string   `json:"websiteUrl"`
	WikipediaURL    string   `json:"wikipediaUrl"`
}

// 神社詳細情報登録画面のリクエストデータ定義
type ShrineContentsRegisterReq struct {
	PlusCode         Field   `json:"plusCode"`
	Furigana         Field   `json:"furigana"`
	AltNames         []Field `json:"altName"`
	Tags             []Field `json:"tag"`
	FoundedYear      Field   `json:"foundedYear"`
	ObjectOfWorships []Field `json:"objectOfWorship"`
	HasGoshuin       Field   `json:"hasGoshuin"`
	WebsiteURL       Field   `json:"websiteUrl"`
	WikipediaURL     Field   `json:"wikipediaUrl"`
}
