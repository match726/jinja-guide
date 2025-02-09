package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
	wikipedia "github.com/match726/jinja-guide/tree/main/server/infrastructure/wikipedia"
)

type ShrineDetailUsecase interface {
	GetShrineDetailByPlusCode(ctx context.Context, plusCode string) (*model.ShrineDetailsResp, error)
}

type shrineDetailUsecase struct {
	sr  repository.ShrineRepository
	scr repository.ShrineContentsRepository
}

func NewShrineDetailUsecase(sr repository.ShrineRepository, scr repository.ShrineContentsRepository) ShrineDetailUsecase {
	return &shrineDetailUsecase{sr: sr, scr: scr}
}

func (sdu shrineDetailUsecase) GetShrineDetailByPlusCode(ctx context.Context, plusCode string) (*model.ShrineDetailsResp, error) {

	var shrs []*model.Shrine
	var shrcs []*model.ShrineContents
	var err error

	query1 := fmt.Sprintf(`SELECT shr.name, shr.address, shr.std_area_code, shr.plus_code, shr.seq, shr.place_id, shr.latitude, shr.longitude, shr.created_at, shr.updated_at
							FROM t_shrines shr
							WHERE shr.plus_code = '%s'`, plusCode)

	shrs, err = sdu.sr.GetShrines(ctx, query1)
	if err != nil {
		return nil, err
	}

	query2 := fmt.Sprintf(`SELECT shrc.id, shrc.seq, shrc.keyword1, COALESCE(shrc.keyword2, '') AS keyword2, shrc.content1, COALESCE(shrc.content2, '') AS content2, COALESCE(shrc.content3, '') AS content3, shrc.created_at, shrc.updated_at
              FROM t_shrine_contents shrc
              WHERE shrc.keyword1 = '%s'
              ORDER BY shrc.id, shrc.seq, shrc.keyword1, shrc.keyword2`, plusCode)

	shrcs, err = sdu.scr.GetShrineContents(ctx, query2)
	if err != nil {
		return nil, err
	}

	// shrdを初期化
	shrd := &model.ShrineDetailsResp{}

	shrd.Name = shrs[0].Name
	shrd.Address = shrs[0].Address
	shrd.PlaceID = shrs[0].PlaceID

	for _, shrc := range shrcs {

		switch shrc.Id {
		case 1:
			// 振り仮名の設定
			shrd.Furigana = shrc.Content1
		case 2:
			// 別名称の設定
			shrd.AltName = append(shrd.AltName, shrc.Content1)
		case 3:
			// 説明の設定
			shrd.Description = shrc.Content1
		case 4:
			// 関連タグの設定
			shrd.Tags = append(shrd.Tags, shrc.Content1)
		case 5:
			// 創建年の設定
			if _, err = strconv.Atoi(shrc.Content1); err == nil {
				shrd.FoundedYear = shrc.Content1 + "年"
			} else {
				shrd.FoundedYear = shrc.Content1
			}
		case 6:
			// 御祭神の設定
			shrd.ObjectOfWorship = append(shrd.ObjectOfWorship, shrc.Content1)
		case 7:
			// 社格の設定
			shrd.ShrineRank = append(shrd.ShrineRank, shrc.Content1)
		case 8:
			//御朱印の設定
			if shrc.Content1 == "あり" {
				shrd.HasGoshuin = true
			}
		case 9:
			// 公式サイトの設定
			shrd.WebsiteURL = shrc.Content1
		case 10:
			// Wikipediaの設定
			shrd.WikipediaURL = shrc.Content1
		}

	}

	// Wikipediaから情報取得
	if len(shrd.WikipediaURL) != 0 {
		image, extract, err := wikipedia.GetShrineDetailsFromWikipedia(shrd.WikipediaURL)
		if err != nil {
			return shrd, fmt.Errorf("%w", err)
		}

		shrd.Image = image
		if len(shrd.Description) == 0 {
			shrd.Description = extract
		}
	}

	if len(shrd.AltName) == 0 {
		shrd.AltName = []string{"登録なし"}
	}
	if len(shrd.Tags) == 0 {
		shrd.Tags = []string{"登録なし"}
	}
	if len(shrd.ObjectOfWorship) == 0 {
		shrd.ObjectOfWorship = []string{"登録なし"}
	}
	if len(shrd.ShrineRank) == 0 {
		shrd.ShrineRank = []string{"登録なし"}
	}

	return shrd, err

}
