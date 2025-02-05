package model

// 標準地域コードテーブルのデータ部の定義(タイムスタンプを含めない)
type StdAreaCodeDataSection struct {
	StdAreaCode     string `json:"stdAreaCode"`
	PrefAreaCode    string `json:"prefAreaCode"`
	SubPrefAreaCode string `json:"subPrefAreaCode"`
	MunicAreaCode1  string `json:"municAreaCode1"`
	MunicAreaCode2  string `json:"municAreaCode2"`
	PrefName        string `json:"prefName"`
	SubPrefName     string `json:"subPrefName"`
	MunicName1      string `json:"municName1"`
	MunicName2      string `json:"municName2"`
}

// 標準地域コード一覧画面のレスポンスデータ定義
type StdAreaCodeListResp struct {
	StdAreaCodeDataSection
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// 当道府県一覧画面のレスポンスデータ定義
type StdAreaCodeRelationshipResp struct {
	StdAreaCode    string `json:"stdAreaCode"`
	Name           string `json:"name"`
	SupStdAreaCode string `json:"supStdAreaCode"`
	Kinds          string `json:"kinds"`
	HasChild       bool   `json:"hasChild"`
}
