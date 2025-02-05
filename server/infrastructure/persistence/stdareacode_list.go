package persistence

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/database"
)

type stdAreaCodeListPersistence struct {
	pg *database.Postgres
}

func NewStdAreaCodeListPersistence(pg *database.Postgres) repository.StdAreaCodeListRepository {
	return &stdAreaCodeListPersistence{pg: pg}
}

func (s *stdAreaCodeListPersistence) GetStdAreaCodeDataSection(ctx context.Context, query string) (sacdss []model.StdAreaCodeDataSection, err error) {

	rows, err := s.pg.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("クエリ実行失敗: %w", err)
	}
	defer rows.Close()

	sacdss, err = pgx.CollectRows(rows, pgx.RowToStructByName[model.StdAreaCodeDataSection])
	if err != nil {
		return nil, fmt.Errorf("コレクト失敗: %w", err)
	}

	return sacdss, nil

}

func (s *stdAreaCodeListPersistence) GetStdAreaCodeList(ctx context.Context, query string) (saclrs []*model.StdAreaCodeListResp, err error) {

	rows, err := s.pg.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("クエリ実行失敗: %w", err)
	}
	defer rows.Close()

	saclrs, err = pgx.CollectRows(rows, pgx.RowToStructByName[*model.StdAreaCodeListResp])
	if err != nil {
		return nil, fmt.Errorf("コレクト失敗: %w", err)
	}

	return saclrs, nil

}
