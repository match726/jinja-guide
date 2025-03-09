package repository

import (
	"context"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
)

type ShrineRegisterRepository interface {
	GetRegisterShrines(ctx context.Context, query string) ([]*model.ShrineRegisterReq, error)
	DeleteRegisterShrine(ctx context.Context, query string) error
}
