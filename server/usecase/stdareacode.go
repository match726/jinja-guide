package usecase

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
)

type StdAreaCodeUsecase interface {
	GetAllStdAreaCodeRelationshipList(ctx context.Context) ([]*model.StdAreaCodeRelationshipResp, error)
}

type stdAreaCodeUsecase struct {
	saclr repository.StdAreaCodeListRepository
}

func NewStdAreaCodeUsecase(saclr repository.StdAreaCodeListRepository) StdAreaCodeUsecase {
	return &stdAreaCodeUsecase{saclr: saclr}
}

func (sacu stdAreaCodeUsecase) GetAllStdAreaCodeRelationshipList(ctx context.Context) (sacrrs []*model.StdAreaCodeRelationshipResp, err error) {

	var sacs []*model.StdAreaCode
	msacrr := make(map[string]model.StdAreaCodeRelationshipResp)

	query := `SELECT shr.std_area_code, sac.pref_area_code, sac.subpref_area_code, sac.munic_area_code1, sac.munic_area_code2, sac.pref_name, sac.subpref_name, sac.munic_name1, sac.munic_name2, sac.created_at, sac.updated_at
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON shr.std_area_code = sac.std_area_code
					GROUP BY shr.std_area_code, sac.pref_area_code, sac.subpref_area_code, sac.munic_area_code1, sac.munic_area_code2, sac.pref_name, sac.subpref_name, sac.munic_name1, sac.munic_name2, sac.created_at, sac.updated_at
					ORDER BY shr.std_area_code`

	sacs, err = sacu.saclr.GetStdAreaCodes(ctx, query)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(sacs); i++ {

		prefCode, _ := strconv.Atoi(sacs[i].StdAreaCode[0:2])
		municCode, _ := strconv.Atoi(sacs[i].StdAreaCode[2:5])

		switch {
		case prefCode == 13 && municCode >= 100 && municCode <= 199:
			// 東京都の特別区部に属する区の場合
			msacrr[sacs[i].PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].PrefAreaCode,
				Name:           sacs[i].PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msacrr[sacs[i].SubPrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].SubPrefAreaCode,
				Name:           sacs[i].SubPrefName,
				SupStdAreaCode: sacs[i].PrefAreaCode,
				Kinds:          "SubPref",
				HasChild:       true,
			}
			msacrr[sacs[i].StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].StdAreaCode,
				Name:           sacs[i].MunicName2,
				SupStdAreaCode: sacs[i].SubPrefAreaCode,
				Kinds:          "Ward",
				HasChild:       false,
			}
		case municCode >= 100 && municCode <= 199:
			// 政令指定都市に属する区の場合
			msacrr[sacs[i].PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].PrefAreaCode,
				Name:           sacs[i].PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msacrr[sacs[i].MunicAreaCode1] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].MunicAreaCode1,
				Name:           sacs[i].MunicName1,
				SupStdAreaCode: sacs[i].PrefAreaCode,
				Kinds:          "City",
				HasChild:       true,
			}
			msacrr[sacs[i].StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].StdAreaCode,
				Name:           sacs[i].MunicName2,
				SupStdAreaCode: sacs[i].MunicAreaCode1,
				Kinds:          "Ward",
				HasChild:       false,
			}
		case municCode >= 201 && municCode <= 299:
			// 政令指定都市以外の市の場合
			msacrr[sacs[i].PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].PrefAreaCode,
				Name:           sacs[i].PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msacrr[sacs[i].StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].StdAreaCode,
				Name:           sacs[i].MunicName1,
				SupStdAreaCode: sacs[i].PrefAreaCode,
				Kinds:          "City",
				HasChild:       false,
			}
		case prefCode == 01 && municCode >= 300:
			// 北海道の振興局に属する町村の場合
			msacrr[sacs[i].PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].PrefAreaCode,
				Name:           sacs[i].PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msacrr[sacs[i].SubPrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].SubPrefAreaCode,
				Name:           sacs[i].SubPrefName,
				SupStdAreaCode: sacs[i].PrefAreaCode,
				Kinds:          "SubPref",
				HasChild:       true,
			}
			msacrr[sacs[i].MunicAreaCode1] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].MunicAreaCode1,
				Name:           sacs[i].MunicName1,
				SupStdAreaCode: sacs[i].SubPrefAreaCode,
				Kinds:          "District",
				HasChild:       true,
			}
			msacrr[sacs[i].StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].StdAreaCode,
				Name:           sacs[i].MunicName2,
				SupStdAreaCode: sacs[i].MunicAreaCode1,
				Kinds:          "Town/Village",
				HasChild:       false,
			}
		case prefCode == 13 && municCode >= 360:
			// 東京都の支庁(離島)に属する町村の場合
			msacrr[sacs[i].PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].PrefAreaCode,
				Name:           sacs[i].PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msacrr[sacs[i].SubPrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].SubPrefAreaCode,
				Name:           sacs[i].SubPrefName,
				SupStdAreaCode: sacs[i].PrefAreaCode,
				Kinds:          "SubPref",
				HasChild:       true,
			}
			msacrr[sacs[i].StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].StdAreaCode,
				Name:           sacs[i].MunicName2,
				SupStdAreaCode: sacs[i].SubPrefAreaCode,
				Kinds:          "Town/Village",
				HasChild:       false,
			}
		case municCode >= 300:
			// 北海道以外の郡に属する町村の場合
			msacrr[sacs[i].PrefAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].PrefAreaCode,
				Name:           sacs[i].PrefName,
				SupStdAreaCode: "",
				Kinds:          "Pref",
				HasChild:       true,
			}
			msacrr[sacs[i].MunicAreaCode1] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].MunicAreaCode1,
				Name:           sacs[i].MunicName1,
				SupStdAreaCode: sacs[i].PrefAreaCode,
				Kinds:          "District",
				HasChild:       true,
			}
			msacrr[sacs[i].StdAreaCode] = model.StdAreaCodeRelationshipResp{
				StdAreaCode:    sacs[i].StdAreaCode,
				Name:           sacs[i].MunicName2,
				SupStdAreaCode: sacs[i].MunicAreaCode1,
				Kinds:          "Town/Village",
				HasChild:       false,
			}
		default:
			// 上記に該当しない場合(スキップ扱いとする)
			fmt.Printf("[エラー] PrefName: %s, SubPrefName: %s, MunicName1: %s, MunicName2: %s, StdAreaCode: %s\n", sacs[i].PrefName, sacs[i].SubPrefName, sacs[i].MunicName1, sacs[i].MunicName2, sacs[i].StdAreaCode)

		}

	}

	// mapのキー(標準地域コード)を元にソートする
	keys := getKeysSacrr(msacrr)
	sort.Strings(keys)
	for _, k := range keys {
		value := msacrr[k]
		sacrrs = append(sacrrs, &value)
	}

	return sacrrs, nil

}

func getKeysSacrr(m map[string]model.StdAreaCodeRelationshipResp) []string {

	keys := []string{}

	for k := range m {
		keys = append(keys, k)
	}

	return keys

}
