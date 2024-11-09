package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	olc "github.com/google/open-location-code/go"
	"github.com/jackc/pgx/v5"
	"googlemaps.github.io/maps"
)

type Shrine struct {
	Name        string
	Address     string
	StdAreaCode string
	PlusCode    string
	Seq         string
	PlaceID     string
	Latitude    float64
	Longitude   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
		return fmt.Errorf("PlaceAPI接続失敗： %s\n", err)
	}

	req := &maps.TextSearchRequest{
		Query:    shr.Name,
		Language: "ja",
	}

	resp, err := client.TextSearch(context.Background(), req)
	if err != nil {
		return fmt.Errorf("PlaceAPI情報取得失敗： %s\n", err)
	} else {
		log.Printf("%#v\n", resp)
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
		"@name":        shr.Name,
		"@address":     shr.Address,
		"@stdAreaCode": shr.StdAreaCode,
		"@plusCode":    shr.PlusCode,
		"@seq":         shr.Seq,
		"@placeId":     shr.PlaceID,
		"@latitude":    shr.Latitude,
		"@longitude":   shr.Longitude,
		"@createdAt":   shr.CreatedAt,
		"@updatedAt":   shr.UpdatedAt,
	}

	_, err := pg.dbPool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("INSERT失敗： %s\n", err)
	}

	return nil

}
