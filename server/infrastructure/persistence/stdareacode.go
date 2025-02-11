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

func NewStdAreaCodePersistence(pg *database.Postgres) repository.StdAreaCodeRepository {
	return &stdAreaCodePersistence{pg: pg}
}

// 標準地域コードテーブル削除
func (s *stdAreaCodePersistence) TruncateStdAreaCode(ctx context.Context) error {

	query := `TRUNCATE TABLE m_stdareacode`

	_, err := s.pg.DbPool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("[標準地域コードテーブル削除失敗]: %w", err)
	}

	return nil

}

// 標準地域コードテーブル登録
func (s *stdAreaCodePersistence) BulkInsertStdAreaCode(ctx context.Context, rows [][]any) (cnt int64, err error) {

	cnt, err = s.pg.DbPool.CopyFrom(
		ctx,
		pgx.Identifier{"m_stdareacode"},
		[]string{"std_area_code", "pref_area_code", "subpref_area_code", "munic_area_code1", "munic_area_code2", "pref_name", "subpref_name", "munic_name1", "munic_name2", "created_at", "updated_at"},
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		return 0, fmt.Errorf("[コピー失敗]: %w", err)
	}

	if int(cnt) != len(rows) {
		return 0, fmt.Errorf("[レコード不一致検知]: 全件数=%d, 登録件数=%d", len(rows), cnt)
	}

	return cnt, nil

}

// 標準地域コードテーブル取得
func (s *stdAreaCodePersistence) GetStdAreaCodes(ctx context.Context, query string) (psacs []*model.StdAreaCode, err error) {

	var sacs []model.StdAreaCode

	rows, err := s.pg.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("[クエリ実行失敗]: %w", err)
	}
	defer rows.Close()

	sacs, err = pgx.CollectRows(rows, pgx.RowToStructByPos[model.StdAreaCode])
	if err != nil {
		return nil, fmt.Errorf("[コレクト失敗]: %w", err)
	}

	for _, sac := range sacs {
		psacs = append(psacs, &sac)
	}

	return psacs, nil

}
