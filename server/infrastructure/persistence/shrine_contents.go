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

// ShrineContents構造体作成
func (s *shrineContentsPersistence) NewShrineContents(id int, seq int, keyword1 string, keyword2 string, content1 string, content2 string, content3 string) *model.ShrineContents {
	return &model.ShrineContents{
		Id:        id,
		Seq:       seq,
		Keyword1:  keyword1,
		Keyword2:  keyword2,
		Content1:  content1,
		Content2:  content2,
		Content3:  content3,
		CreatedAt: database.GetNowTime(),
		UpdatedAt: database.GetNowTime(),
	}
}

// 神社詳細情報テーブル登録
func (s *shrineContentsPersistence) InsertShrineContents(ctx context.Context, shrc *model.ShrineContents) (err error) {

	query := `INSERT INTO t_shrine_contents (
						id,
						seq,
						keyword1,
						keyword2,
						content1,
						content2,
						content3,
						created_at,
						updated_at
						) VALUES (
						@id,
						@seq,
						@keyword1,
						@keyword2,
						@content1,
						@content2,
						@content3,
						@createdAt,
						@updatedAt
						)`

	args := pgx.NamedArgs{
		"id":        shrc.Id,
		"seq":       shrc.Seq,
		"keyword1":  shrc.Keyword1,
		"keyword2":  shrc.Keyword2,
		"content1":  shrc.Content1,
		"content2":  shrc.Content2,
		"content3":  shrc.Content3,
		"createdAt": database.GetNowTime(),
		"updatedAt": database.GetNowTime(),
	}

	_, err = s.pg.DbPool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("[神社詳細情報テーブル登録失敗]: %w", err)
	}

	return nil

}

// 神社詳細情報テーブル取得
func (s *shrineContentsPersistence) GetShrineContents(ctx context.Context, query string) (pshrcs []*model.ShrineContents, err error) {

	var shrcs []model.ShrineContents

	rows, err := s.pg.DbPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("[クエリ実行失敗]: %w", err)
	}
	defer rows.Close()

	shrcs, err = pgx.CollectRows(rows, pgx.RowToStructByPos[model.ShrineContents])
	if err != nil {
		return nil, fmt.Errorf("[コレクト失敗]: %w", err)
	}

	//fmt.Printf("shrcs: %v\n", shrcs)

	for _, shrc := range shrcs {
		pshrcs = append(pshrcs, &shrc)
	}

	return pshrcs, nil

}

// 神社詳細情報テーブル登録用のSeq取得
func (s *shrineContentsPersistence) GetShrineContentsNextSeq(ctx context.Context, shrc *model.ShrineContents) (err error) {

	var seq int

	query := fmt.Sprintf(`SELECT COALESCE(MAX(shrc.seq), 0)
	FROM t_shrine_contents shrc
	WHERE shrc.keyword1 = '%s'
	AND shrc.id = %d`, shrc.Keyword1, shrc.Id)

	err = s.pg.DbPool.QueryRow(ctx, query).Scan(&seq)
	if err != nil {
		return fmt.Errorf("[SEQ取得失敗]: %w", err)
	}

	// ShrineContents構造体のSeqへセット
	shrc.Seq = seq + 1

	return nil

}
