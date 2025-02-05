package usecase

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
)

type StdAreaCodeListUsecase interface {
	GetAllStdAreaCodeRelationshipList(ctx context.Context) ([]*model.StdAreaCodeRelationshipResp, error)
}

type stdAreaCodeListUsecase struct {
	saclr repository.StdAreaCodeListRepository
}

func NewStdAreaCodeListUsecase(saclr repository.StdAreaCodeListRepository) StdAreaCodeListUsecase {
	return &stdAreaCodeListUsecase{saclr: saclr}
}

func (slu stdAreaCodeListUsecase) GetAllStdAreaCodeRelationshipList(ctx context.Context) (saclrs []*model.StdAreaCodeRelationshipResp, err error) {

	var sacdss []model.StdAreaCodeDataSection
	msh := make(map[string]model.StdAreaCodeRelationshipResp)

	query := `SELECT shr.std_area_code, sac.pref_area_code, sac.subpref_area_code, sac.munic_area_code1, sac.munic_area_code2, sac.pref_name, sac.subpref_name, sac.munic_name1, sac.munic_name2
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON shr.std_area_code = sac.std_area_code
					GROUP BY shr.std_area_code, sac.pref_area_code, sac.subpref_area_code, sac.munic_area_code1, sac.munic_area_code2, sac.pref_name, sac.subpref_name, sac.munic_name1, sac.munic_name2
					ORDER BY shr.std_area_code`

	sacdss, err = slu.saclr.GetStdAreaCodeDataSection(ctx, query)
	if err != nil {
		return nil, err
	}

	for _, sacds := range sacdss {

		prefCode, _ := strconv.Atoi(sacds.StdAreaCode[0:2])
		municCode, _ := strconv.Atoi(sacds.StdAreaCode[2:5])

		switch {
		case prefCode == 13 && municCode >= 100 && municCode <= 199:
			// 東京都の特別区部に属する区の場合
			msh[sacds.PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.PrefAreaCode,
				Name:           sacds.PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msh[sacds.SubPrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.SubPrefAreaCode,
				Name:           sacds.SubPrefName,
				SupStdAreaCode: sacds.PrefAreaCode,
				Kinds:          "SubPref",
				HasChild:       true,
			}
			msh[sacds.StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.StdAreaCode,
				Name:           sacds.MunicName2,
				SupStdAreaCode: sacds.SubPrefAreaCode,
				Kinds:          "Ward",
				HasChild:       false,
			}
		case municCode >= 100 && municCode <= 199:
			// 政令指定都市に属する区の場合
			msh[sacds.PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.PrefAreaCode,
				Name:           sacds.PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msh[sacds.MunicAreaCode1] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.MunicAreaCode1,
				Name:           sacds.MunicName1,
				SupStdAreaCode: sacds.PrefAreaCode,
				Kinds:          "City",
				HasChild:       true,
			}
			msh[sacds.StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.StdAreaCode,
				Name:           sacds.MunicName2,
				SupStdAreaCode: sacds.MunicAreaCode1,
				Kinds:          "Ward",
				HasChild:       false,
			}
		case municCode >= 201 && municCode <= 299:
			// 政令指定都市以外の市の場合
			msh[sacds.PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.PrefAreaCode,
				Name:           sacds.PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msh[sacds.StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.StdAreaCode,
				Name:           sacds.MunicName1,
				SupStdAreaCode: sacds.PrefAreaCode,
				Kinds:          "City",
				HasChild:       false,
			}
		case prefCode == 01 && municCode >= 300:
			// 北海道の振興局に属する町村の場合
			msh[sacds.PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.PrefAreaCode,
				Name:           sacds.PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msh[sacds.SubPrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.SubPrefAreaCode,
				Name:           sacds.SubPrefName,
				SupStdAreaCode: sacds.PrefAreaCode,
				Kinds:          "SubPref",
				HasChild:       true,
			}
			msh[sacds.MunicAreaCode1] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.MunicAreaCode1,
				Name:           sacds.MunicName1,
				SupStdAreaCode: sacds.SubPrefAreaCode,
				Kinds:          "District",
				HasChild:       true,
			}
			msh[sacds.StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.StdAreaCode,
				Name:           sacds.MunicName2,
				SupStdAreaCode: sacds.MunicAreaCode1,
				Kinds:          "Town/Village",
				HasChild:       false,
			}
		case prefCode == 13 && municCode >= 360:
			// 東京都の支庁(離島)に属する町村の場合
			msh[sacds.PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.PrefAreaCode,
				Name:           sacds.PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msh[sacds.SubPrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.SubPrefAreaCode,
				Name:           sacds.SubPrefName,
				SupStdAreaCode: sacds.PrefAreaCode,
				Kinds:          "SubPref",
				HasChild:       true,
			}
			msh[sacds.StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.StdAreaCode,
				Name:           sacds.MunicName2,
				SupStdAreaCode: sacds.SubPrefAreaCode,
				Kinds:          "Town/Village",
				HasChild:       false,
			}
		case municCode >= 300:
			// 北海道以外の郡に属する町村の場合
			msh[sacds.PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.PrefAreaCode,
				Name:           sacds.PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msh[sacds.MunicAreaCode1] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.MunicAreaCode1,
				Name:           sacds.MunicName1,
				SupStdAreaCode: sacds.PrefAreaCode,
				Kinds:          "District",
				HasChild:       true,
			}
			msh[sacds.StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacds.StdAreaCode,
				Name:           sacds.MunicName2,
				SupStdAreaCode: sacds.MunicAreaCode1,
				Kinds:          "Town/Village",
				HasChild:       false,
			}
		default:
			// 上記に該当しない場合(スキップ扱いとする)
			fmt.Printf("[エラー] PrefName: %s, SubPrefName: %s, MunicName1: %s, MunicName2: %s, StdAreaCode: %s\n", sacds.PrefName, sacds.SubPrefName, sacds.MunicName1, sacds.MunicName2, sacds.StdAreaCode)
		}

	}

	// mapのキー(標準地域コード)を元にソートする
	keys := getStdAreaCodeRelationshipRespKeys(msh)
	sort.Strings(keys)
	for _, k := range keys {
		value := msh[k]
		saclrs = append(saclrs, &value)
	}

	return saclrs, nil

}

func getStdAreaCodeRelationshipRespKeys(m map[string]model.StdAreaCodeRelationshipResp) []string {

	keys := []string{}

	for k := range m {
		keys = append(keys, k)
	}

	return keys

}
