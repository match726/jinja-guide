package models

import (
	"github.com/rs/xid"
)

// XIDの取得
func GetXID() (id string) {

	id = xid.New().String()
	return id

}
