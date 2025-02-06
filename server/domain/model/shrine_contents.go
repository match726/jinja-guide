package model

import "time"

// 神社詳細テーブル定義
type ShrineContents struct {
	Id        int       `json:"id"`
	Seq       int       `json:"seq"`
	Keyword1  string    `json:"keyword1"`
	Keyword2  string    `json:"keyword2"`
	Content1  string    `json:"content1"`
	Content2  string    `json:"content2"`
	Content3  string    `json:"content3"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
