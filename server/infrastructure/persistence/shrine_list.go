package persistence

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/database"
)

type shrineListPersistence struct {
	pg *database.Postgres
}

func NewShrineListPersistence(pg *database.Postgres) repository.ShrineListRepository {
	return &shrineListPersistence{pg: pg}
}

func (s *shrineListPersistence) GetShrineListResps(ctx context.Context, query string) (pslrsps []*model.ShrineListResp, err error) {

	var slrsps []model.ShrineListResp

	rows, err := s.pg.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("クエリ実行失敗: %w", err)
	}
	defer rows.Close()

	slrsps, err = pgx.CollectRows(rows, pgx.RowToStructByPos[model.ShrineListResp])
	if err != nil {
		return nil, fmt.Errorf("コレクト失敗: %w", err)
	}

	for _, slrsp := range slrsps {
		pslrsps = append(pslrsps, &slrsp)
	}

	return pslrsps, nil

}
