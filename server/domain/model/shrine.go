package model

import "time"

// 神社テーブル定義
type Shrine struct {
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	StdAreaCode string    `json:"stdAreaCode"`
	PlusCode    string    `json:"plusCode"`
	Seq         string    `json:"seq"`
	PlaceID     string    `json:"placeId"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
