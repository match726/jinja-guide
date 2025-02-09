package persistence

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/database"
)

type shrineContentsPersistence struct {
	pg *database.Postgres
}

func NewShrineContentsPersistence(pg *database.Postgres) repository.ShrineContentsRepository {
	return &shrineContentsPersistence{pg: pg}
}

func (s *shrineContentsPersistence) GetShrineContents(ctx context.Context, query string) (pshrcs []*model.ShrineContents, err error) {

	var shrcs []model.ShrineContents

	rows, err := s.pg.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("クエリ実行失敗: %w", err)
	}
	defer rows.Close()

	shrcs, err = pgx.CollectRows(rows, pgx.RowToStructByPos[model.ShrineContents])
	if err != nil {
		return nil, fmt.Errorf("コレクト失敗: %w", err)
	}

	fmt.Printf("shrcs: %v\n", shrcs)

	for _, shrc := range shrcs {
		pshrcs = append(pshrcs, &shrc)
	}

	return pshrcs, nil

}
