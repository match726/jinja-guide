package persistence

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/database"
)

type shrinePersistence struct {
	pg *database.Postgres
}

func NewShrinePersistence(pg *database.Postgres) repository.ShrineRepository {
	return &shrinePersistence{pg: pg}
}

// 神社テーブル登録
func (s *shrinePersistence) InsertShrine(ctx context.Context, shr *model.Shrine) error {

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
		"createdAt":   database.GetNowTime(),
		"updatedAt":   database.GetNowTime(),
	}

	_, err := s.pg.DbPool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("神社テーブル登録失敗: %w", err)
	}

	return nil

}

// 神社テーブル取得
func (s *shrinePersistence) GetShrines(ctx context.Context, query string) (pshrs []*model.Shrine, err error) {

	var shrs []model.Shrine

	rows, err := s.pg.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("クエリ実行失敗: %w", err)
	}
	defer rows.Close()

	shrs, err = pgx.CollectRows(rows, pgx.RowToStructByPos[model.Shrine])
	if err != nil {
		return nil, fmt.Errorf("コレクト失敗: %w", err)
	}

	for _, shr := range shrs {
		pshrs = append(pshrs, &shr)
	}

	return pshrs, nil

}
