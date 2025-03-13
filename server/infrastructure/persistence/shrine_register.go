package persistence

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/database"
)

type shrineRegisterPersistence struct {
	pg *database.Postgres
}

func NewShrineRegisterPersistence(pg *database.Postgres) repository.ShrineRegisterRepository {
	return &shrineRegisterPersistence{pg: pg}
}

func (r *shrineRegisterPersistence) GetRegisterShrines(ctx context.Context, query string) (pshrrs []*model.ShrineRegister, err error) {

	var shrrs []model.ShrineRegister

	rows, err := r.pg.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("[クエリ実行失敗]: %w", err)
	}
	defer rows.Close()

	shrrs, err = pgx.CollectRows(rows, pgx.RowToStructByPos[model.ShrineRegister])
	if err != nil {
		return nil, fmt.Errorf("[コレクト失敗]: %w", err)
	}

	for _, shrr := range shrrs {
		pshrrs = append(pshrrs, &shrr)
	}

	return pshrrs, nil

}

func (r *shrineRegisterPersistence) DeleteRegisterShrine(ctx context.Context, query string) (err error) {

	_, err = r.pg.DbPool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("[クエリ実行失敗]: %w", err)
	}

	return nil

}
