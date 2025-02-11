package repository

import (
	"context"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
)

type StdAreaCodeRepository interface {
	TruncateStdAreaCode(ctx context.Context) error
	BulkInsertStdAreaCode(ctx context.Context, rows [][]any) (cnt int64, err error)
	GetStdAreaCodes(ctx context.Context, query string) ([]*model.StdAreaCode, error)
}
