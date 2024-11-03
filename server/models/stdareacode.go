package models

import (
	"fmt"
	"os"
	"time"
)

// ★北海道の郡のコード付番をstd_area_codeのINSERT時に行う
// ⇒神社の一覧取得時に引っ掛からないため

type StdAreaCode struct {
	StdAreaCode     string
	PrefAreaCode    string
	SubPrefAreaCode string
	MunicAreaCode1  string
	MunicAreaCode2  string
	PrefName        string
	SubPrefName     string
	MunicName1      string
	MunicName2      string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type StdAreaCodes []StdAreaCode

// e-Statの統計LODから最新の標準地域コードを取得する
func GetAllStdAreaCodesFromEstat() (sacs StdAreaCodes) {

	var current time.Time
	current = time.Now()

	prefix := `PREFIX sacs: <http://data.e-stat.go.jp/lod/terms/sacs#>
PREFIX dcterms: <http://purl.org/dc/terms/>
PREFIX ic: <http://imi.go.jp/ns/core/rdf#>
PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>`

	// 県の抽出
	query1 := `SELECT DISTINCT ?areacode ?psac ?spsac ?m1sac ?m2sac ?pref ?subpref ?munic1 ?munic2
WHERE {
	{
	?s a sacs:StandardAreaCode ;
		dcterms:identifier ?areacode ;
		dcterms:identifier ?psac ;
		sacs:prefectureLabel ?pref ;
		sacs:administrativeClass ?adclass .
	FILTER ( ?adclass = sacs:Prefecture )
	}`

	/*
		// 振興局・支庁の抽出
		query2 := `  UNION
		{
		?s a sacs:StandardAreaCode ;
			dcterms:identifier ?areacode ;
			dcterms:isPartOf / dcterms:identifier ?psac ;
			dcterms:identifier ?spsac ;
			sacs:prefectureLabel ?pref ;
			ic:表記 ?subpref ;
			sacs:administrativeClass ?adclass .
		FILTER ( lang(?subpref) = "ja" )
		FILTER ( ?adclass = sacs:SubPrefecture )
		}`

		// 特別区部(東京都)の抽出
		query3 := `  UNION
		{
		?s a sacs:StandardAreaCode ;
			dcterms:identifier ?areacode ;
			dcterms:isPartOf / dcterms:identifier ?psac ;
			dcterms:identifier ?spsac ;
			sacs:prefectureLabel ?pref ;
			ic:表記 ?subpref ;
			sacs:administrativeClass ?adclass .
		FILTER ( lang(?subpref) = "ja" )
		FILTER ( ?adclass = sacs:SpecialWardsArea )
		}`

		// 市の抽出
		query4 := `  UNION
		{
		?s a sacs:StandardAreaCode ;
			dcterms:identifier ?areacode ;
			dcterms:isPartOf / dcterms:identifier ?psac ;
			dcterms:identifier ?m1sac ;
			sacs:prefectureLabel ?pref ;
			ic:表記 ?munic1 ;
			sacs:administrativeClass ?adclass .
		FILTER ( lang(?munic1) = "ja" )
		FILTER ( ?adclass IN( sacs:City, sacs:CoreCity, sacs:DesignatedCity, sacs:SpecialCity ) )
		}`

		// 郡の抽出
		query5 := `  UNION
		{
		?s a sacs:StandardAreaCode ;
			dcterms:identifier ?areacode ;
			dcterms:isPartOf / dcterms:identifier ?psac ;
			dcterms:identifier ?m1sac ;
			sacs:prefectureLabel ?pref ;
			ic:表記 ?munic1 ;
			sacs:administrativeClass ?adclass ;
		FILTER ( lang(?munic1) = "ja" )
		FILTER ( ?adclass = sacs:District )
		}`

		// 振興局・支庁に属する、北海道の町村の抽出 (北海道の振興局)
		query6 := `  UNION
		{
		?s a sacs:StandardAreaCode ;
			dcterms:identifier ?areacode ;
			dcterms:isPartOf / dcterms:isPartOf / dcterms:identifier ?psac ;
			dcterms:isPartOf / dcterms:identifier ?spsac ;
			dcterms:identifier ?m2sac ;
			sacs:prefectureLabel ?pref ;
			dcterms:isPartOf ?spo ;
			sacs:districtOfSubPrefecture ?munic1 ;
			ic:表記 ?munic2 ;
			sacs:administrativeClass ?adclass1 ;
			dcterms:isPartOf / sacs:administrativeClass ?adclass2 .
		?spo ic:表記 ?subpref .
		FILTER REGEX ( ?pref, "北海道")
		FILTER ( lang(?subpref) = "ja" )
		FILTER ( lang(?munic1) = "ja" )
		FILTER ( lang(?munic2) = "ja" )
		FILTER ( ?adclass1 IN( sacs:Town, sacs:Village, sacs:Ward ) )
		FILTER ( ?adclass2 = sacs:SubPrefecture )
		MINUS { ?spo dcterms:valid ?spo2 }
		}`

		// 振興局・支庁に属する、東京の町村の抽出 (東京の離島)
		query7 := `  UNION
		{
		?s a sacs:StandardAreaCode ;
			dcterms:identifier ?areacode ;
			dcterms:isPartOf / dcterms:isPartOf / dcterms:identifier ?psac ;
			dcterms:isPartOf / dcterms:identifier ?spsac ;
			dcterms:identifier ?m2sac ;
			sacs:prefectureLabel ?pref ;
			dcterms:isPartOf ?spo ;
			ic:表記 ?munic2 ;
			sacs:administrativeClass ?adclass1 ;
			dcterms:isPartOf / sacs:administrativeClass ?adclass2 .
		?spo ic:表記 ?subpref .
		FILTER REGEX ( ?pref, "東京")
		FILTER ( lang(?subpref) = "ja" )
		FILTER ( lang(?munic2) = "ja" )
		FILTER ( ?adclass1 IN( sacs:Town, sacs:Village, sacs:Ward ) )
		FILTER ( ?adclass2 = sacs:SubPrefecture )
		MINUS { ?spo dcterms:valid ?spo2 }
		}`

		// 振興局・支庁に属さない、区町村の抽出
		query8 := `  UNION
		{
		?s a sacs:StandardAreaCode ;
			dcterms:identifier ?areacode ;
			dcterms:isPartOf / dcterms:isPartOf / dcterms:identifier ?psac ;
			dcterms:isPartOf / dcterms:identifier ?m1sac ;
			dcterms:identifier ?m2sac ;
			sacs:prefectureLabel ?pref ;
			dcterms:isPartOf / ic:表記 ?munic1 ;
			ic:表記 ?munic2 ;
			sacs:administrativeClass ?adclass1 ;
			dcterms:isPartOf / sacs:administrativeClass ?adclass2 .
		FILTER ( lang(?munic1) = "ja" )
		FILTER ( ?adclass1 IN( sacs:Town, sacs:Village, sacs:Ward ) )
		FILTER ( ?adclass2 != sacs:SubPrefecture )
		}`

		// 東京２３区の抽出
		query9 := `  UNION
		{
		?s a sacs:StandardAreaCode ;
			dcterms:identifier ?areacode ;
			dcterms:isPartOf / dcterms:isPartOf / dcterms:identifier ?psac ;
			dcterms:isPartOf / dcterms:identifier ?spsac ;
			dcterms:identifier ?m2sac ;
			sacs:prefectureLabel ?pref ;
			dcterms:isPartOf / ic:表記 ?subpref ;
			ic:表記 ?munic2 ;
			sacs:administrativeClass ?adclass .
		FILTER ( lang(?subpref) = "ja" )
		FILTER ( lang(?munic2) = "ja" )
		FILTER ( ?adclass = sacs:SpecialWard )
		}`
	*/

	query10 := `  MINUS { ?s dcterms:valid ?o }
}
ORDER BY ?areacode`

	query := prefix + query1 + query10
	//query := prefix + query1 + query2 + query3 + query4 + query5 + query6 + query7 + query8 + query9 + query10

	fmt.Println("SPARQL発行前")
	resp := QuerySparql(os.Getenv("ESTAT_ENDPOINT"), query)
	fmt.Println("SPARQL発行後")

	for _, m := range resp.Results.Bindings {
		sacs = append(sacs, StdAreaCode{m["AREACODE"].Value, m["PSAC"].Value, m["SPSAC"].Value, m["M1SAC"].Value, m["M2SAC"].Value, m["PREF"].Value, m["SUBPREF"].Value, m["MUNIC1"].Value, m["MUNIC2"].Value, current, current})
	}

	return sacs

}
