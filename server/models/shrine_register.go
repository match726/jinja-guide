package models

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"regexp"

// 	"github.com/jackc/pgx/v5"
// 	"googlemaps.github.io/maps"
// )

// // 神社詳細テーブルへの登録
// // ★重複時の制御が必要
// func (pg *Postgres) InsertShrineContents(ctx context.Context, id int, content string, plusCode string, seqHandler int) (err error) {

// 	var seq int = 1

// 	// 登録するSEQを取得する
// 	if seqHandler == 1 {
// 		seq, err = pg.GetShrineContentsSeq(ctx, id, plusCode)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	query := `INSERT INTO t_shrine_contents (
// 						id,
// 						seq,
// 						keyword1,
// 						keyword2,
// 						content1,
// 						content2,
// 						content3,
// 						created_at,
// 						updated_at
// 						) VALUES (
// 						@id,
// 						@seq,
// 						@keyword1,
// 						@keyword2,
// 						@content1,
// 						@content2,
// 						@content3,
// 						@createdAt,
// 						@updatedAt
// 						)`

// 	args := pgx.NamedArgs{
// 		"id":        id,
// 		"seq":       seq,
// 		"keyword1":  plusCode,
// 		"keyword2":  "",
// 		"content1":  content,
// 		"content2":  "",
// 		"content3":  "",
// 		"createdAt": GetNowTime(),
// 		"updatedAt": GetNowTime(),
// 	}

// 	_, err = pg.dbPool.Exec(ctx, query, args)
// 	if err != nil {
// 		return fmt.Errorf("INSERT失敗： %w", err)
// 	}

// 	return nil

// }
