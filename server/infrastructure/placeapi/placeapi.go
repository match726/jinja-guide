package placeapi

import (
	"context"
	"fmt"
	"os"

	"googlemaps.github.io/maps"
)

// PlaceAPI(Google)から情報を取得する
func QueryPlaceAPI(ctx context.Context, keyword1 string, keyword2 string) (resp maps.PlacesSearchResponse, err error) {

	apikey := os.Getenv("GOOGLE_PLACE_API_KEY")

	client, err := maps.NewClient(maps.WithAPIKey(apikey))
	if err != nil {
		return maps.PlacesSearchResponse{}, fmt.Errorf("PlaceAPI接続失敗: %w", err)
	}

	req := &maps.TextSearchRequest{
		Query:    keyword1 + "　" + keyword2,
		Language: "ja",
	}

	resp, err = client.TextSearch(ctx, req)
	if err != nil {
		return maps.PlacesSearchResponse{}, fmt.Errorf("PlaceAPI情報取得失敗: %w", err)
	}

	return resp, nil

}
