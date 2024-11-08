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

	query := `TRUNCATE TABLE m_stdareacode`

	_, err = pg.dbPool.Exec(context.Background(), query)

	if err != nil {
		return fmt.Errorf("標準地域コード TRUNCATE失敗： %s\n", err)
	}

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
			case municCode >= 510 && municCode < 540:
				// 宗谷総合振興局の場合
				switch m.MunicName1 {
				case "宗谷郡":
					m.MunicAreaCode1 = "0151A"
				case "枝幸郡":
					m.MunicAreaCode1 = "0151B"
				case "天塩郡":
					m.MunicAreaCode1 = "0151C"
				case "礼文郡":
					m.MunicAreaCode1 = "0151D"
				case "利尻郡":
					m.MunicAreaCode1 = "0151E"
				}
			case municCode >= 540 && municCode < 570:
				// オホーツク総合振興局の場合
				switch m.MunicName1 {
				case "網走郡":
					m.MunicAreaCode1 = "0154A"
				case "斜里郡":
					m.MunicAreaCode1 = "0154B"
				case "常呂郡":
					m.MunicAreaCode1 = "0154C"
				case "紋別郡":
					m.MunicAreaCode1 = "0154D"
				}
			case municCode >= 570 && municCode < 600:
				// 胆振総合振興局の場合
				switch m.MunicName1 {
				case "虻田郡":
					m.MunicAreaCode1 = "0157A"
				case "有珠郡":
					m.MunicAreaCode1 = "0157B"
				case "白老郡":
					m.MunicAreaCode1 = "0157C"
				case "勇払郡":
					m.MunicAreaCode1 = "0157D"
				}
			case municCode >= 600 && municCode < 630:
				// 日高振興局の場合
				switch m.MunicName1 {
				case "沙流郡":
					m.MunicAreaCode1 = "0160A"
				case "新冠郡":
					m.MunicAreaCode1 = "0160B"
				case "浦河郡":
					m.MunicAreaCode1 = "0160C"
				case "様似郡":
					m.MunicAreaCode1 = "0160D"
				case "幌泉郡":
					m.MunicAreaCode1 = "0160E"
				case "日高郡":
					m.MunicAreaCode1 = "0160F"
				}
			case municCode >= 630 && municCode < 660:
				// 十勝総合振興局の場合
				switch m.MunicName1 {
				case "河東郡":
					m.MunicAreaCode1 = "0163A"
				case "上川郡":
					m.MunicAreaCode1 = "0163B"
				case "河西郡":
					m.MunicAreaCode1 = "0163C"
				case "広尾郡":
					m.MunicAreaCode1 = "0163D"
				case "中川郡":
					m.MunicAreaCode1 = "0163E"
				case "足寄郡":
					m.MunicAreaCode1 = "0163F"
				case "十勝郡":
					m.MunicAreaCode1 = "0163G"
				}
			case municCode >= 660 && municCode < 690:
				// 釧路総合振興局の場合
				switch m.MunicName1 {
				case "釧路郡":
					m.MunicAreaCode1 = "0166A"
				case "厚岸郡":
					m.MunicAreaCode1 = "0166B"
				case "川上郡":
					m.MunicAreaCode1 = "0166C"
				case "阿寒郡":
					m.MunicAreaCode1 = "0166D"
				case "白糠郡":
					m.MunicAreaCode1 = "0166E"
				}
			case municCode >= 690:
				// 根室振興局の場合
				switch m.MunicName1 {
				case "野付郡":
					m.MunicAreaCode1 = "0169A"
				case "標津郡":
					m.MunicAreaCode1 = "0169B"
				case "目梨郡":
					m.MunicAreaCode1 = "0169C"
				case "色丹郡":
					m.MunicAreaCode1 = "0169D"
				case "国後郡":
					m.MunicAreaCode1 = "0169E"
				case "択捉郡":
					m.MunicAreaCode1 = "0169F"
				case "紗那郡":
					m.MunicAreaCode1 = "0169G"
				case "蘂取郡":
					m.MunicAreaCode1 = "0169H"
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
		return fmt.Errorf("標準地域コード 書き込み失敗： %s\n", err)
	} else {
		fmt.Printf("UpdateStdAreaCode: 書き込み成功(%d行)\n", cnt)
	}

	if int(cnt) != len(sacs) {
		return fmt.Errorf("標準地域コード レコード不一致： %d\n", cnt)
	}

	return err

}

// e-Statの統計LODから最新の標準地域コードを取得する
func GetAllStdAreaCodesFromEstat() (sacs []StdAreaCode) {

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

	query10 := `  MINUS { ?s dcterms:valid ?o }
}
ORDER BY ?areacode`

	query := prefix + query1 + query2 + query3 + query4 + query5 + query6 + query7 + query8 + query9 + query10

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
		return nil, fmt.Errorf("標準地域コード 取得失敗： %w\n", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[StdAreaCodeGet])

}

// 特定の都道府県に属する標準地域コードの一覧を取得する (神社の住所からの標準地域コード取得用)
func (pg *Postgres) GetStdAreaCodeListByPrefName(prefname string) (sacs []StdAreaCode, err error) {

	query := fmt.Sprintf(`SELECT std_area_code, pref_name, subpref_name, munic_name1, munic_name2
					FROM m_stdareacode
					WHERE pref_name = '%s'`, prefname)

	rows, err := pg.dbPool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("標準地域コード一覧 取得失敗： %w\n", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[StdAreaCode])

}
