package models

import (
	"time"

	"github.com/rs/xid"
)

// 現在日時の取得
func GetNowTime() (current time.Time) {

	jstZone := time.FixedZone("Asia/Tokyo", 9*60*60)
	current = time.Now().In(jstZone)

	return current

}

// XIDの取得
func GetXID() (id string) {

	id = xid.New().String()
	return id

}
