package repository

import (
	"context"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
)

type StdAreaCodeRepository interface {
	GetStdAreaCodes(ctx context.Context, query string) ([]model.StdAreaCode, error)
}
