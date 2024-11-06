package models

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
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
}

// SELECT用の構造体(タイムスタンプをYYYY/MM/DD HH24:MM:SS形式で取得する)
type StdAreaCodeGet struct {
	StdAreaCode
	CreatedAt string
	UpdatedAt string
}

// 標準エリアコードの最新化
func (pg *Postgres) UpdateStdAreaCode() (err error) {

	var sacs []StdAreaCode
	var rows [][]interface{}

	/*
		query := `TRUNCATE TABLE m_stdareacode`

		_, err = pg.dbPool.Exec(context.Background(), query)

		if err != nil {
			return fmt.Errorf("標準地域コード TRUNCATE失敗： %s", err)
		}
	*/

	sacs = GetAllStdAreaCodesFromEstat()
	current := GetNowTime()

	for _, m := range sacs {

		prefCode, _ := strconv.Atoi(m.StdAreaCode[0:2])
		municCode, _ := strconv.Atoi(m.StdAreaCode[2:5])
		// 北海道の振興局に属する町村の場合
		// 北海道の郡には標準地域コードが振られていないため、付番する
		if prefCode == 01 && municCode >= 300 {
			switch {
			case municCode >= 300 && municCode < 330:
				// 石狩振興局の場合
				switch m.MunicName1 {
				case "石狩郡":
					m.MunicAreaCode1 = "1300A"
				}
			case municCode >= 330 && municCode < 360:
				//  渡島総合振興局の場合
				switch m.MunicName1 {
				case "松前郡":
					m.MunicAreaCode1 = "0133A"
				case "上磯郡":
					m.MunicAreaCode1 = "0133B"
				case "亀田郡":
					m.MunicAreaCode1 = "0133C"
				case "茅部郡":
					m.MunicAreaCode1 = "0133D"
				case "二海郡":
					m.MunicAreaCode1 = "0133E"
				case "山越郡":
					m.MunicAreaCode1 = "0133F"
				}
			case municCode >= 360 && municCode < 390:
				// 檜山振興局の場合
				switch m.MunicName1 {
				case "檜山郡":
					m.MunicAreaCode1 = "0136A"
				case "爾志郡":
					m.MunicAreaCode1 = "0136B"
				case "奥尻郡":
					m.MunicAreaCode1 = "0136C"
				case "瀬棚郡":
					m.MunicAreaCode1 = "0136D"
				case "久遠郡":
					m.MunicAreaCode1 = "0136E"
				}
			case municCode >= 390 && municCode < 420:
				// 後志総合振興局の場合
				switch m.MunicName1 {
				case "島牧郡":
					m.MunicAreaCode1 = "0139A"
				case "寿都郡":
					m.MunicAreaCode1 = "0139B"
				case "磯谷郡":
					m.MunicAreaCode1 = "0139C"
				case "虻田郡":
					m.MunicAreaCode1 = "0139D"
				case "岩内郡":
					m.MunicAreaCode1 = "0139E"
				case "古宇郡":
					m.MunicAreaCode1 = "0139F"
				case "積丹郡":
					m.MunicAreaCode1 = "0139G"
				case "古平郡":
					m.MunicAreaCode1 = "0139H"
				case "余市郡":
					m.MunicAreaCode1 = "0139I"
				}
			case municCode >= 420 && municCode < 450:
				// 空知総合振興局の場合
				switch m.MunicName1 {
				case "空知郡":
					m.MunicAreaCode1 = "0142A"
				case "夕張郡":
					m.MunicAreaCode1 = "0142B"
				case "樺戸郡":
					m.MunicAreaCode1 = "0142C"
				case "雨竜郡":
					m.MunicAreaCode1 = "0142D"
				}
			case municCode >= 450 && municCode < 480:
				// 上川総合振興局の場合
				switch m.MunicName1 {
				case "上川郡":
					m.MunicAreaCode1 = "0145A"
				case "空知郡":
					m.MunicAreaCode1 = "0145B"
				case "勇払郡":
					m.MunicAreaCode1 = "0145C"
				case "中川郡":
					m.MunicAreaCode1 = "0145D"
				case "雨竜郡":
					m.MunicAreaCode1 = "0145E"
				}
			case municCode >= 480 && municCode < 510:
				// 留萌振興局の場合
				switch m.MunicName1 {
				case "増毛郡":
					m.MunicAreaCode1 = "0148A"
				case "留萌郡":
					m.MunicAreaCode1 = "0148B"
				case "苫前郡":
					m.MunicAreaCode1 = "0148C"
				case "天塩郡":
					m.MunicAreaCode1 = "0148D"
				}
			}
		}
		rows = append(rows, []interface{}{m.StdAreaCode, m.PrefAreaCode, m.SubPrefAreaCode, m.MunicAreaCode1, m.MunicAreaCode2, m.PrefName, m.SubPrefName, m.MunicName1, m.MunicName2, current, current})
	}

	cnt, err := pg.dbPool.CopyFrom(
		context.Background(),
		pgx.Identifier{"m_stdareacode"},
		[]string{"std_area_code", "pref_area_code", "subpref_area_code", "munic_area_code1", "munic_area_code2", "pref_name", "subpref_name", "munic_name1", "munic_name2", "created_at", "updated_at"},
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		return fmt.Errorf("標準地域コード 書き込み失敗： %s", err)
	} else {
		fmt.Printf("UpdateStdAreaCode: 書き込み成功(%d行)\n", cnt)
	}

	if int(cnt) != len(sacs) {
		return fmt.Errorf("標準地域コード レコード不一致： %d", cnt)
	}

	return err

}

// e-Statの統計LODから最新の標準地域コードを取得する
func GetAllStdAreaCodesFromEstat() (sacs []StdAreaCode) {

	prefix := `PREFIX sacs: <http://data.e-stat.go.jp/lod/terms/sacs#>
PREFIX dcterms: <http://purl.org/dc/terms/>
PREFIX ic: <http://imi.go.jp/ns/core/rdf#>
PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>`
	/*
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
	*/

	// 振興局・支庁に属する、北海道の町村の抽出 (北海道の振興局)
	//query6 := `  UNION
	query6 := `SELECT DISTINCT ?areacode ?psac ?spsac ?m1sac ?m2sac ?pref ?subpref ?munic1 ?munic2
		WHERE {
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
	/*
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

	query := prefix + query6 + query7 + query10

	resp := QuerySparql(os.Getenv("ESTAT_ENDPOINT"), query)

	for _, m := range resp.Results.Bindings {
		sacs = append(sacs, StdAreaCode{m["AREACODE"].Value, m["PSAC"].Value, m["SPSAC"].Value, m["M1SAC"].Value, m["M2SAC"].Value, m["PREF"].Value, m["SUBPREF"].Value, m["MUNIC1"].Value, m["MUNIC2"].Value})
	}

	return sacs

}

func (pg *Postgres) GetStdAreaCodes() ([]StdAreaCodeGet, error) {

	query := `SELECT std_area_code, pref_area_code, subpref_area_code, munic_area_code1, munic_area_code2, pref_name, subpref_name, munic_name1, munic_name2, to_char(created_at,'YYYY/MM/DD HH24:MI:SS') AS "created_at", to_char(updated_at,'YYYY/MM/DD HH24:MI:SS') AS "updated_at"
					FROM m_stdareacode
					ORDER BY std_area_code`

	rows, err := pg.dbPool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("標準地域コード 取得失敗： %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[StdAreaCodeGet])

}
