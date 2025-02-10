package repository

import (
	"context"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
)

type ShrineContentsRepository interface {
	NewShrineContents(id int, seq int, keyword1 string, keyword2 string, content1 string, content2 string, content3 string) *model.ShrineContents
	InsertShrineContents(ctx context.Context, shrc *model.ShrineContents) error
	GetShrineContents(ctx context.Context, query string) ([]*model.ShrineContents, error)
	GetShrineContentsNextSeq(ctx context.Context, shrc *model.ShrineContents) error
}
