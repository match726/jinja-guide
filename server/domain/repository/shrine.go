package repository

import (
	"context"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
)

type ShrineRepository interface {
	GetShrines(ctx context.Context, query string) ([]*model.Shrine, error)
	InsertShrine(ctx context.Context, shr *model.Shrine) error
}
