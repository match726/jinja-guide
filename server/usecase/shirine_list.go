package usecase

import (
	"context"
	"fmt"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
)

type ShrineListUsecase interface {
	GetShrineListByStdAreaCode(ctx context.Context, kinds string, stdAreaCode string) ([]model.ShrineListResp, error)
	GetShrineListByTag(ctx context.Context, tag string) ([]model.ShrineListResp, error)
}

type shrineListUsecase struct {
	slr repository.ShrineListRepository
}

func NewShrineListUsecase(slr repository.ShrineListRepository) ShrineListUsecase {
	return &shrineListUsecase{slr: slr}
}

func (slu shrineListUsecase) GetShrineListByStdAreaCode(ctx context.Context, kinds string, stdAreaCode string) ([]model.ShrineListResp, error) {

	var query string

	switch kinds {
	case "Pref":
		query = fmt.Sprintf(`SELECT shr.name, shr.address, shr.plus_code, shr.place_id, ARRAY[]::VARCHAR[] AS object_of_worship, CASE shrc.content1 WHEN 'あり' THEN true ELSE false END AS has_goshuin
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.pref_area_code = '%s'
						AND shr.std_area_code = sac.std_area_code
					LEFT JOIN t_shrine_contents shrc
						ON shrc.id = 8
						AND shr.plus_code = shrc.keyword1
					ORDER BY shr.std_area_code, shr.address, shr.name`, stdAreaCode)
	case "SubPref":
		query = fmt.Sprintf(`SELECT shr.name, shr.address, shr.plus_code, shr.place_id, ARRAY[]::VARCHAR[] AS object_of_worship, CASE shrc.content1 WHEN 'あり' THEN true ELSE false END AS has_goshuin
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.subpref_area_code = '%s'
						AND shr.std_area_code = sac.std_area_code
					LEFT JOIN t_shrine_contents shrc
						ON shrc.id = 8
						AND shr.plus_code = shrc.keyword1
					ORDER BY shr.std_area_code, shr.address, shr.name`, stdAreaCode)
	case "City", "District":
		query = fmt.Sprintf(`SELECT shr.name, shr.address, shr.plus_code, shr.place_id, ARRAY[]::VARCHAR[] AS object_of_worship, CASE shrc.content1 WHEN 'あり' THEN true ELSE false END AS has_goshuin
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.munic_area_code1 = '%s'
						AND shr.std_area_code = sac.std_area_code
					LEFT JOIN t_shrine_contents shrc
						ON shrc.id = 8
						AND shr.plus_code = shrc.keyword1
					ORDER BY shr.std_area_code, shr.address, shr.name`, stdAreaCode)
	case "Town/Village", "Ward":
		query = fmt.Sprintf(`SELECT shr.name, shr.address, shr.plus_code, shr.place_id, ARRAY[]::VARCHAR[] AS object_of_worship, CASE shrc.content1 WHEN 'あり' THEN true ELSE false END AS has_goshuin
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.munic_area_code2 = '%s'
						AND shr.std_area_code = sac.std_area_code
					LEFT JOIN t_shrine_contents shrc
						ON shrc.id = 8
						AND shr.plus_code = shrc.keyword1
					ORDER BY shr.std_area_code, shr.address, shr.name`, stdAreaCode)
	}

	shrls, err := slu.slr.GetShrineList(ctx, query)
	if err != nil {
		return nil, err
	}

	return shrls, nil

}

func (slu shrineListUsecase) GetShrineListByTag(ctx context.Context, tag string) ([]model.ShrineListResp, error) {

	query := fmt.Sprintf(`SELECT shr.name, shr.address, shr.plus_code, shr.place_id, ARRAY[]::VARCHAR[] AS object_of_worship, CASE shrc2.content1 WHEN 'あり' THEN true ELSE false END AS has_goshuin
						FROM t_shrine_contents shrc
						INNER JOIN t_shrines shr
							ON shrc.keyword1 = shr.plus_code
						LEFT JOIN t_shrine_contents shrc2
							ON shrc2.id = 8
							AND shr.plus_code = shrc2.keyword1
						WHERE shrc.id = 4
							AND shrc.content1 = '%s'
						ORDER BY shr.std_area_code, shr.address, shr.name`, tag)

	shrls, err := slu.slr.GetShrineList(ctx, query)
	if err != nil {
		return nil, err
	}

	return shrls, nil

}
