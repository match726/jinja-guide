package models

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"time"

	olc "github.com/google/open-location-code/go"
	"github.com/jackc/pgx/v5"
	"googlemaps.github.io/maps"
)

type Shrine struct {
	name        string
	address     string
	stdAreaCode string
	plusCode    string
	seq         string
	placeId     string
	latitude    float64
	longitude   float64
	createdAt   time.Time
	updatedAt   time.Time
}

type Shrinecontents struct {
	id        int
	keyword1  string
	keyword2  string
	content1  string
	content2  string
	content3  string
	createdAt time.Time
	updatedAt time.Time
}

type ShrineDetails struct {
	name            string
	furigana        string
	image           string
	altName         []string
	address         string
	placeId         string
	description     string
	tags            []string
	foundedYear     string
	objectOfWorship []string
	shrineRank      []string
	hasGoshuin      bool
	websiteUrl      string
	wikipediaUrl    string
}

// 住所から都道府県部分を抽出する
func ExtractPrefname(address string) string {

	// ★県の一致が\Wでは雑過ぎない??
	reg, _ := regexp.Compile(`^東京都|^北海道|^(大阪|京都)府|^\W{2,3}県`)
	pref := reg.FindString(address)

	return pref

}

// PlaceAPI(Google)から情報を取得する
func GetLocnInfoFromPlaceAPI(shr *Shrine) error {

	apikey := os.Getenv("GOOGLE_PLACE_API_KEY")

	client, err := maps.NewClient(maps.WithAPIKey(apikey))
	if err != nil {
		return fmt.Errorf("PlaceAPI接続失敗： %w", err)
	}

	req := &maps.TextSearchRequest{
		Query:    shr.name + "　" + shr.address,
		Language: "ja",
	}

	resp, err := client.TextSearch(context.Background(), req)
	if err != nil {
		return fmt.Errorf("PlaceAPI情報取得失敗： %w", err)
	} else {
		fmt.Printf("%#v\n", resp)
	}

	shr.placeId = resp.Results[0].PlaceID
	shr.latitude = resp.Results[0].Geometry.Location.Lat
	shr.longitude = resp.Results[0].Geometry.Location.Lng
	shr.plusCode = olc.Encode(shr.latitude, shr.longitude, 11)

	return nil

}

// 神社情報の登録
// ★重複時の制御が必要
func (pg *Postgres) InsertShrine(shr *Shrine) error {

	query := `INSERT INTO t_shrines (
						name,
						address,
						std_area_code,
						plus_code,
						seq,
						place_id,
						latitude,
						longitude,
						created_at,
						updated_at
						) VALUES (
						@name,
						@address,
						@stdAreaCode,
						@plusCode,
						@seq,
						@placeId,
						@latitude,
						@longitude,
						@createdAt,
						@updatedAt
						)`

	args := pgx.NamedArgs{
		"name":        shr.name,
		"address":     shr.address,
		"stdAreaCode": shr.stdAreaCode,
		"plusCode":    shr.plusCode,
		"seq":         shr.seq,
		"placeId":     shr.placeId,
		"latitude":    shr.latitude,
		"longitude":   shr.longitude,
		"createdAt":   GetNowTime(),
		"updatedAt":   GetNowTime(),
	}

	_, err := pg.dbPool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("INSERT失敗： %w", err)
	}

	return nil

}

func (pg *Postgres) GetShrinesBystdAreaCode(sacr *SacRelationship) (shrs []*Shrine, err error) {

	var query string

	switch sacr.kinds {
	case "Pref":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdAreaCode sac
						ON sac.pref_area_code = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "SubPref":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdAreaCode sac
						ON sac.subpref_area_code = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "City", "District":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdAreaCode sac
						ON sac.munic_area_code1 = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "Town/Village", "Ward":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdAreaCode sac
						ON sac.munic_area_code2 = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	}

	rows, err := pg.dbPool.Query(context.Background(), query, sacr.stdAreaCode)
	if err != nil {
		return nil, fmt.Errorf("神社一覧 取得失敗： %w", err)
	}
	defer rows.Close()

	for rows.Next() {

		var shr Shrine

		err = rows.Scan(&shr.name, &shr.address, &shr.plusCode, &shr.placeId)
		if err != nil {
			return nil, fmt.Errorf("スキャン失敗： %w", err)
		}

		shrs = append(shrs, &shr)

	}

	return shrs, err

}

// 神社の詳細情報を取得
func (pg *Postgres) GetShrineDetails(shr *Shrine) (shrd ShrineDetails, err error) {

	query1 := `SELECT shr.name, shr.address, shr.place_id
						FROM t_shrines shr
						WHERE shr.plus_code = $1`

	row := pg.dbPool.QueryRow(context.Background(), query1, shr.plusCode)

	err = row.Scan(&shrd.name, &shrd.address, &shrd.placeId)
	if err != nil {
		return shrd, fmt.Errorf("スキャン１失敗： %w", err)
	}

	query2 := `SELECT shrc.id, shrc.content1, shrc.content2, shrc.content3
              FROM t_shrine_contents shrc
              WHERE shrc.keyword1 = $1
              ORDER BY shrc.id, shrc.keyword1, shrc.keyword2`

	rows, err := pg.dbPool.Query(context.Background(), query2, shr.plusCode)
	if err != nil {
		return shrd, fmt.Errorf("神社詳細情報 取得失敗： %w", err)
	}
	defer rows.Close()

	for rows.Next() {

		var shrc Shrinecontents

		err = rows.Scan(&shrc.id, &shrc.content1, &shrc.content2, &shrc.content3)
		if err != nil {
			return shrd, fmt.Errorf("スキャン失敗： %w", err)
		}

		switch shrc.id {
		case 1:
			// 振り仮名の設定
			shrd.furigana = shrc.content1
		case 2:
			// 別名称の設定
			shrd.altName = append(shrd.altName, shrc.content1)
		case 3:
			// 説明の設定
			shrd.description = shrc.content1
		case 4:
			// 関連タグの設定
			shrd.tags = append(shrd.tags, shrc.content1)
		case 5:
			// 創建年の設定
			shrd.foundedYear = shrc.content1
		case 6:
			// 御祭神の設定
			shrd.objectOfWorship = append(shrd.objectOfWorship, shrc.content1)
		case 7:
			// 社格の設定
			shrd.shrineRank = append(shrd.shrineRank, shrc.content1)
		//case 8:
		// 御朱印の設定
		case 9:
			// 公式サイトの設定
			shrd.websiteUrl = shrc.content1
		case 10:
			// Wikipediaの設定
			shrd.wikipediaUrl = shrc.content1
		}

	}

	if len(shrd.altName) == 0 {
		shrd.altName = []string{"登録なし"}
	}
	if len(shrd.tags) == 0 {
		shrd.tags = []string{"登録なし"}
	}
	if len(shrd.objectOfWorship) == 0 {
		shrd.objectOfWorship = []string{"登録なし"}
	}
	if len(shrd.shrineRank) == 0 {
		shrd.shrineRank = []string{"登録なし"}
	}

	return shrd, err

}
