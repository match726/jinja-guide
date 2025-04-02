package usecase

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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

	// 神社テーブルから基本情報を取得
	query1 := fmt.Sprintf(`SELECT shr.name, shr.address, shr.std_area_code, shr.plus_code, shr.seq, shr.place_id, shr.latitude, shr.longitude, shr.created_at, shr.updated_at
							FROM t_shrines shr
							WHERE shr.plus_code = '%s'`, plusCode)

	shrs, err = sdu.sr.GetShrines(ctx, query1)
	if err != nil {
		return nil, err
	}

	// 神社詳細情報テーブルから詳細情報を取得
	query2 := fmt.Sprintf(`SELECT shrc.id, shrc.seq, shrc.keyword1, COALESCE(shrc.keyword2, '') AS keyword2, shrc.content1, COALESCE(shrc.content2, '') AS content2, COALESCE(shrc.content3, '') AS content3, shrc.created_at, shrc.updated_at
              FROM t_shrine_contents shrc
              WHERE shrc.keyword1 = '%s'
              ORDER BY shrc.id, shrc.seq, shrc.keyword1, shrc.keyword2`, plusCode)

	shrcs, err = sdu.scr.GetShrineContents(ctx, query2)
	if err != nil {
		return nil, err
	}

	// ShrineDetailsResp構造体を初期化
	shrd := &model.ShrineDetailsResp{}

	// 神社テーブルからShrineDetailsResp構造体に項目設定
	shrd.Name = shrs[0].Name
	shrd.Address = shrs[0].Address
	shrd.PlaceID = shrs[0].PlaceID

	// 神社詳細情報テーブルからShrineDetailsResp構造体に項目設定
	for _, shrc := range shrcs {

		fieldResp := model.Field{Id: 1, Seq: strconv.Itoa(shrc.Seq), Content1: shrc.Content1, Content2: shrc.Content2, Content3: shrc.Content3}

		switch shrc.Id {
		case 1:
			// 振り仮名の設定
			shrd.Furigana = fieldResp
		case 2:
			// 別名称の設定
			shrd.AltName = append(shrd.AltName, fieldResp)
		case 3:
			// 説明の設定
			shrd.Description = fieldResp
		case 4:
			// 関連タグの設定
			shrd.Tags = append(shrd.Tags, fieldResp)
		case 5:
			// 創建年の設定
			shrd.FoundedYear = fieldResp
			if _, err = strconv.Atoi(shrc.Content1); err == nil {
				shrd.FoundedYear.Content1 = shrc.Content1 + "年"
			}
		case 6:
			// 御祭神の設定
			shrd.ObjectOfWorship = append(shrd.ObjectOfWorship, fieldResp)
		case 7:
			// 社格の設定
			shrd.ShrineRank = append(shrd.ShrineRank, fieldResp)
		case 8:
			//御朱印の設定
			shrd.HasGoshuin = fieldResp
		case 9:
			// 公式サイトの設定
			shrd.WebsiteURL = fieldResp
		case 10:
			// Wikipediaの設定
			shrd.WikipediaURL = fieldResp
		}

	}

	// Wikipediaから情報取得し、ShrineDetailsResp構造体に設定
	if len(shrd.WikipediaURL.Content1) != 0 {
		title := shrd.WikipediaURL.Content1[strings.LastIndex(shrd.WikipediaURL.Content1, "/")+1:]
		image, extract, err := wikipedia.QueryWikipedia(title)
		if err != nil {
			return shrd, fmt.Errorf("%w", err)
		}

		shrd.Image = image
		if len(shrd.Description.Content1) == 0 {
			shrd.Description.Content1 = extract
		}
	}

	noFieldResp := model.Field{Id: 1, Seq: "", Content1: "登録なし", Content2: "", Content3: ""}

	if len(shrd.AltName) == 0 {
		shrd.AltName = append(shrd.AltName, noFieldResp)
	}
	if len(shrd.Tags) == 0 {
		shrd.Tags = append(shrd.Tags, noFieldResp)
	}
	if len(shrd.ObjectOfWorship) == 0 {
		shrd.ObjectOfWorship = append(shrd.ObjectOfWorship, noFieldResp)
	}
	if len(shrd.ShrineRank) == 0 {
		shrd.ShrineRank = append(shrd.ShrineRank, noFieldResp)
	}

	return shrd, err

}
