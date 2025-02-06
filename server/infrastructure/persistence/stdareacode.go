package persistence

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/database"
)

type stdAreaCodePersistence struct {
	pg *database.Postgres
}

func NewStdAreaCodePersistence(pg *database.Postgres) repository.StdAreaCodeListRepository {
	return &stdAreaCodePersistence{pg: pg}
}

func (s *stdAreaCodePersistence) GetStdAreaCodes(ctx context.Context, query string) (sacs []*model.StdAreaCode, err error) {

	rows, err := s.pg.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("クエリ実行失敗: %w", err)
	}
	defer rows.Close()

	sacs, err = pgx.CollectRows(rows, pgx.RowToStructByPos[*model.StdAreaCode])
	if err != nil {
		return nil, fmt.Errorf("コレクト失敗: %w", err)
	}

	return sacs, nil

}
