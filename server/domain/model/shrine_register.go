package model

type ShrineRegisterReq struct {
	Name         string `json:"name"`
	Furigana     string `json:"furigana"`
	Address      string `json:"address"`
	WikipediaURL string `json:"wikipediaUrl"`
}

type ShrineContentsRegisterReq struct {
	PlusCode        string `json:"plusCode"`
	Furigana        string `json:"furigana"`
	AltName         string `json:"altName"`
	Tags            string `json:"tags"`
	FoundedYear     string `json:"foundedYear"`
	ObjectOfWorship string `json:"objectOfWorship"`
	HasGoshuin      string `json:"hasGoshuin"`
	WebsiteURL      string `json:"websiteUrl"`
	WikipediaURL    string `json:"wikipediaUrl"`
}
