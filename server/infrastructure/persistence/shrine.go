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

	fmt.Printf("shrs: %v\n", shrs)
	for _, shr := range shrs {
		pshrs = append(pshrs, &shr)
	}

	return pshrs, nil

}
