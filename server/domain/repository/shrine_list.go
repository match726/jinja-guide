package repository

import (
	"context"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
)

type ShrineListRepository interface {
	GetShrineList(ctx context.Context, query string) ([]model.ShrineListResp, error)
}
