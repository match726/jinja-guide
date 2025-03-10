package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
	wikipedia "github.com/match726/jinja-guide/tree/main/server/infrastructure/wikipedia"
)

type HomeContentsUsecase interface {
	GetHomeContents(ctx context.Context) (*model.HomeContentsResp, error)
	GetRandomShrines(ctx context.Context) ([]*model.RandomShrines, error)
	GetAllTags(ctx context.Context) ([]*model.AllTags, error)
}

type homeContentsUsecase struct {
	sr  repository.ShrineRepository
	scr repository.ShrineContentsRepository
}

func NewHomeContentsUsecase(sr repository.ShrineRepository, scr repository.ShrineContentsRepository) HomeContentsUsecase {
	return &homeContentsUsecase{sr: sr, scr: scr}
}

func (hcu homeContentsUsecase) GetHomeContents(ctx context.Context) (*model.HomeContentsResp, error) {

	var hcr model.HomeContentsResp
	var rshrs []*model.RandomShrines
	var tags []*model.AllTags
	var err error

	hcr = model.HomeContentsResp{}

	rshrs, err = hcu.GetRandomShrines(ctx)
	if err != nil {
		return nil, err
	}

	hcr.Shrines = append(hcr.Shrines, rshrs...)

	tags, err = hcu.GetAllTags(ctx)
	if err != nil {
		return nil, err
	}

	hcr.Tags = append(hcr.Tags, tags...)

	return &hcr, err

}

func (hcu homeContentsUsecase) GetRandomShrines(ctx context.Context) ([]*model.RandomShrines, error) {

	var shrs []*model.Shrine
	var rshrs []*model.RandomShrines
	var err error

	rshrs = []*model.RandomShrines{}

	// 神社テーブル取得（ランダム3件）
	query1 := `SELECT * FROM t_shrines
						ORDER BY RANDOM() LIMIT 3`

	shrs, err = hcu.sr.GetShrines(ctx, query1)
	if err != nil {
		return nil, err
	}

	for _, shr := range shrs {

		var rshr model.RandomShrines
		var shrcs []*model.ShrineContents
		var wikipediaURL string

		rshr.Name = shr.Name
		rshr.Address = shr.Address
		rshr.PlusCode = shr.PlusCode
		rshr.PlaceId = shr.PlaceID

		// 神社詳細情報テーブルから詳細情報を取得
		query2 := fmt.Sprintf(`SELECT shrc.id, shrc.seq, shrc.keyword1, COALESCE(shrc.keyword2, '') AS keyword2, shrc.content1, COALESCE(shrc.content2, '') AS content2, COALESCE(shrc.content3, '') AS content3, shrc.created_at, shrc.updated_at
								FROM t_shrine_contents shrc
								WHERE shrc.id IN (1, 3, 6, 10)
								AND shrc.keyword1 = '%s'
								ORDER BY shrc.id, shrc.seq, shrc.keyword1, shrc.keyword2`, rshr.PlusCode)

		shrcs, err = hcu.scr.GetShrineContents(ctx, query2)
		if err != nil {
			return nil, err
		}

		for _, shrc := range shrcs {

			switch shrc.Id {
			case 1:
				// 振り仮名の設定
				rshr.Furigana = shrc.Content1
			case 3:
				// 説明の設定
				rshr.Description = shrc.Content1
			case 6:
				// 御祭神の設定
				rshr.ObjectOfWorship = append(rshr.ObjectOfWorship, shrc.Content1)
			case 10:
				// Wikipediaの設定
				wikipediaURL = shrc.Content1
			}

		}

		// Wikipediaから情報取得し、HomeContentsResp構造体に設定
		if len(wikipediaURL) != 0 {
			title := wikipediaURL[strings.LastIndex(wikipediaURL, "/")+1:]
			_, extract, err := wikipedia.QueryWikipedia(title)
			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}

			if len(rshr.Description) == 0 {
				rshr.Description = extract
			}
		}

		if len(rshr.ObjectOfWorship) == 0 {
			rshr.ObjectOfWorship = []string{"登録なし"}
		}
		if len(rshr.Description) == 0 {
			rshr.Description = "説明文なし"
		}

		rshrs = append(rshrs, &rshr)
		fmt.Println(rshrs)
	}

	return rshrs, err

}

func (hcu homeContentsUsecase) GetAllTags(ctx context.Context) ([]*model.AllTags, error) {

	var tags []*model.AllTags
	var err error

	// 神社詳細テーブル取得（関連ワード一覧取得）
	tags, err = hcu.scr.GetTagsWithCount(ctx)
	if err != nil {
		return nil, err
	}

	return tags, err

}
