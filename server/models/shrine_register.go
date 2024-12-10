package models

import (
	"context"
	"fmt"
	"os"
	"regexp"

	olc "github.com/google/open-location-code/go"
	"github.com/jackc/pgx/v5"
	"googlemaps.github.io/maps"
)

// 住所から都道府県部分を抽出する
func ExtractPrefName(address string) string {

	// ★県の一致が\Wでは雑過ぎない??
	reg, _ := regexp.Compile(`^東京都|^北海道|^(大阪|京都)府|^\W{2,3}県`)
	pref := reg.FindString(address)

	return pref

}

// PlaceAPI(Google)から神社の情報を取得する
func GetLocnInfoFromPlaceAPI(ctx context.Context, shr *Shrine) error {

	apikey := os.Getenv("GOOGLE_PLACE_API_KEY")

	client, err := maps.NewClient(maps.WithAPIKey(apikey))
	if err != nil {
		return fmt.Errorf("PlaceAPI接続失敗： %w", err)
	}

	req := &maps.TextSearchRequest{
		Query:    shr.Name + "　" + shr.Address,
		Language: "ja",
	}

	resp, err := client.TextSearch(ctx, req)
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

// 神社テーブルへの登録
// ★重複時の制御が必要
func (pg *Postgres) InsertShrine(ctx context.Context, shr *Shrine) error {

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

	_, err := pg.dbPool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("INSERT失敗： %w", err)
	}

	return nil

}

// 神社詳細テーブルへの登録
// ★重複時の制御が必要
func (pg *Postgres) InsertShrineContents(ctx context.Context, id int, content string, plusCode string) error {

	query := `INSERT INTO t_shrine_contents (
						id,
						seq,
						keyword1,
						keyword2,
						content1,
						content2,
						content3,
						created_at,
						updated_at
						) VALUES (
						@id,
						@seq,
						@keyword1,
						@keyword2,
						@content1,
						@content2,
						@content3,
						@createdAt,
						@updatedAt
						)`

	args := pgx.NamedArgs{
		"id":        id,
		"seq":       1,
		"keyword1":  plusCode,
		"keyword2":  "",
		"content1":  content,
		"content2":  "",
		"content3":  "",
		"createdAt": GetNowTime(),
		"updatedAt": GetNowTime(),
	}

	_, err := pg.dbPool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("INSERT失敗： %w", err)
	}

	return nil

}
