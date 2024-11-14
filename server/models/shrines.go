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
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	StdAreaCode string    `json:"std_area_code"`
	PlusCode    string    `json:"plus_code"`
	Seq         string    `json:"seq"`
	PlaceID     string    `json:"place_id"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ShrineDetails struct {
	Name            string   `json:"name"`
	Furigana        string   `json:"furigana"`
	AltName         string   `json:"alt_name"`
	Address         string   `json:"address"`
	Image           string   `json:"image"`
	Description     string   `json:"description"`
	Tags            []string `json:"tags"`
	FoundedYear     string   `json:"founded_year"`
	ObjectOfWorship []string `json:"object_of_worship"`
	ShrineRank      []string `json:"shrine_rank"`
	HasGoshuin      bool     `json:"has_goshuin"`
	WebsiteURL      string   `json:"website_url"`
	WikipediaURL    string   `json:"wikipedia_url"`
}

// 住所から都道府県部分を抽出する
func ExtractPrefName(address string) string {

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
		Query:    shr.Name,
		Language: "ja",
	}

	resp, err := client.TextSearch(context.Background(), req)
	if err != nil {
		return fmt.Errorf("PlaceAPI情報取得失敗： %w", err)
	} else {
		fmt.Printf("%#v\n", resp)
	}

	shr.PlaceID = resp.Results[0].PlaceID
	shr.Latitude = resp.Results[0].Geometry.Location.Lat
	shr.Longitude = resp.Results[0].Geometry.Location.Lng
	shr.PlusCode = olc.Encode(shr.Latitude, shr.Longitude, 11)

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
		"name":        shr.Name,
		"address":     shr.Address,
		"stdAreaCode": shr.StdAreaCode,
		"plusCode":    shr.PlusCode,
		"seq":         shr.Seq,
		"placeId":     shr.PlaceID,
		"latitude":    shr.Latitude,
		"longitude":   shr.Longitude,
		"createdAt":   GetNowTime(),
		"updatedAt":   GetNowTime(),
	}

	_, err := pg.dbPool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("INSERT失敗： %w", err)
	}

	return nil

}

func (pg *Postgres) GetShrinesByStdAreaCode(sacr *SacRelationship) (shrs []*Shrine, err error) {

	var query string

	switch sacr.Kinds {
	case "Pref":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.pref_area_code = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "SubPref":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.subpref_area_code = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "City", "District":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.munic_area_code1 = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "Town/Village", "Ward":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.munic_area_code2 = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	}

	rows, err := pg.dbPool.Query(context.Background(), query, sacr.StdAreaCode)
	if err != nil {
		return nil, fmt.Errorf("神社一覧 取得失敗： %w", err)
	}
	defer rows.Close()

	for rows.Next() {

		var shr Shrine

		err = rows.Scan(&shr.Name, &shr.Address, &shr.PlusCode, &shr.PlaceID)
		if err != nil {
			return nil, fmt.Errorf("スキャン失敗： %w", err)
		}

		shrs = append(shrs, &shr)

	}

	return shrs, err

}

func (pg *Postgres) GetShrineDetails(shr *Shrine) (shrd ShrineDetails, err error) {

	shrd.Tags = []string{}
	shrd.ObjectOfWorship = []string{}
	shrd.ShrineRank = []string{}

	query := `SELECT shr.name, shr.address
						FROM t_shrines shr
						WHERE shr.plus_code = $1`

	row, err := pg.dbPool.Query(context.Background(), query, shr.PlusCode)
	if err != nil {
		return shrd, fmt.Errorf("神社詳細 取得失敗： %w", err)
	}
	defer row.Close()

	shrd, err = pgx.CollectOneRow(row, pgx.RowToStructByName[ShrineDetails])
	if err != nil {
		return shrd, fmt.Errorf("スキャン失敗： %w", err)
	}

	return shrd, err

}
