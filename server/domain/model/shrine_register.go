package model

type ShrineRegisterReq struct {
	Name         string `json:"name"`
	Furigana     string `json:"furigana"`
	Address      string `json:"address"`
	WikipediaURL string `json:"wikipediaUrl"`
}
