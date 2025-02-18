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

func (r *shrineRegisterPersistence) GetRegisterShrines(ctx context.Context, query string) (pshrrreqs []*model.ShrineRegisterReq, err error) {

	var shrrreqs []model.ShrineRegisterReq

	rows, err := r.pg.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("[クエリ実行失敗]: %w", err)
	}
	defer rows.Close()

	shrrreqs, err = pgx.CollectRows(rows, pgx.RowToStructByPos[model.ShrineRegisterReq])
	if err != nil {
		return nil, fmt.Errorf("[コレクト失敗]: %w", err)
	}

	for _, shrrreq := range shrrreqs {
		pshrrreqs = append(pshrrreqs, &shrrreq)
	}

	return pshrrreqs, nil

}
