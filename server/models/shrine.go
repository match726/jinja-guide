package models

// import (
// 	"context"
// 	"time"

// 	"github.com/match726/jinja-guide/tree/main/server/logger"
// )

// // PlusCodeから神社の登録の有無を判定
// func (pg *Postgres) ExistsShrineByPlusCode(ctx context.Context, plusCode string) bool {

// 	var shr Shrine

// 	query := `SELECT shr.name
// 						FROM t_shrines shr
// 						WHERE shr.plus_code = $1`

// 	err := pg.dbPool.QueryRow(ctx, query, plusCode).Scan(&shr.Name)

// 	if err != nil {
// 		logger.Error(ctx, "SELECT失敗")
// 		return false
// 	}

// 	if len(shr.Name) != 0 {
// 		return true
// 	}

// 	return false

// }
