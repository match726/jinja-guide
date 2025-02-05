package repository

import (
	"context"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
)

type StdAreaCodeListRepository interface {
	GetStdAreaCodeDataSection(ctx context.Context, query string) ([]*model.StdAreaCodeDataSection, error)
	GetStdAreaCodeList(ctx context.Context, query string) ([]*model.StdAreaCodeListResp, error)
}
