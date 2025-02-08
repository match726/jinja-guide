package repository

import (
	"context"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
)

type ShrineContentsRepository interface {
	GetShrineContents(ctx context.Context, query string) ([]*model.ShrineContents, error)
}
