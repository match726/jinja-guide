package models

import (
	"time"
)

// 神社テーブル
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

// 神社詳細テーブル
type ShrineContents struct {
	Id        int
	Seq       int
	Keyword1  string
	Keyword2  string
	Content1  string
	Content2  string
	Content3  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ShrineDetails struct {
	Name            string   `json:"name"`
	Furigana        string   `json:"furigana"`
	Image           string   `json:"image"`
	AltName         []string `json:"altName"`
	Address         string   `json:"address"`
	PlaceID         string   `json:"placeId"`
	Description     string   `json:"description"`
	Tags            []string `json:"tags"`
	FoundedYear     string   `json:"foundedYear"`
	ObjectOfWorship []string `json:"objectOfWorship"`
	ShrineRank      []string `json:"shrineRank"`
	HasGoshuin      bool     `json:"hasGoshuin"`
	WebsiteURL      string   `json:"websiteUrl"`
	WikipediaURL    string   `json:"wikipediaUrl"`
}

func (shr *Shrine) ShrineName(name string) {
	shr.Name = name
}

func (shr *Shrine) ShrineAddress(address string) {
	shr.Address = address
}
