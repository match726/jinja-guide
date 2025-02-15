package usecase

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	olc "github.com/google/open-location-code/go"
	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/placeapi"
)

type ShrineRegisterUsecase interface {
	GetStdAreaCodeByAddress(ctx context.Context, shrrreq *model.ShrineRegisterReq) (sac string, err error)
	GetLocnInfoFromPlaceAPI(ctx context.Context, shrrreq *model.ShrineRegisterReq, sac string) (shr *model.Shrine, err error)
	RegisterShrine(ctx context.Context, shr *model.Shrine) (err error)
	RegisterShrineContents(ctx context.Context, id int, seq int, keyword1 string, keyword2 string, content1 string, content2 string, content3 string, seqHandler int) (err error)
	ExistsShrineByPlusCode(ctx context.Context, plusCode string) bool
}

type shrineRegisterUsecase struct {
	sacr repository.StdAreaCodeRepository
	sr   repository.ShrineRepository
	scr  repository.ShrineContentsRepository
}

func NewShrineRegisterUsecase(sacr repository.StdAreaCodeRepository, sr repository.ShrineRepository, scr repository.ShrineContentsRepository) ShrineRegisterUsecase {
	return &shrineRegisterUsecase{sacr: sacr, sr: sr, scr: scr}
}

// 住所から該当する標準地域コードを取得
func (sru shrineRegisterUsecase) GetStdAreaCodeByAddress(ctx context.Context, shrrreq *model.ShrineRegisterReq) (sac string, err error) {

	var sacs []*model.StdAreaCode

	// 住所から都道府県を取得
	reg, _ := regexp.Compile(`^東京都|^北海道|^(大阪|京都)府|^\W{2,3}県`)
	pref := reg.FindString(shrrreq.Address)

	// 該当の都道府県の標準地域コード一覧を取得
	query := fmt.Sprintf(`SELECT sac.std_area_code, sac.pref_area_code, sac.subpref_area_code, sac.munic_area_code1, sac.munic_area_code2, sac.pref_name, sac.subpref_name, sac.munic_name1, sac.munic_name2, sac.created_at, sac.updated_at
					FROM m_stdareacode sac
					WHERE sac.pref_name = '%s'`, pref)

	sacs, err = sru.sacr.GetStdAreaCodes(ctx, query)
	if err != nil {
		return "", err
	}

	// 住所に当てはまる標準地域コードを紐付け
	for i := len(sacs) - 1; i >= 0; i-- {
		if sacs[i].MunicName1 == "" && sacs[i].MunicName2 == "" {
			continue
		} else {
			keyword := sacs[i].PrefName + sacs[i].MunicName1 + sacs[i].MunicName2
			if strings.HasPrefix(shrrreq.Address, keyword) {
				sac = sacs[i].StdAreaCode
				break
			}
		}
	}

	return sac, nil

}

// PlaceAPIから位置情報(PlaceID、緯度、経度)とPlusCodeを取得
// ⇒Shrine構造体の形式で返す
func (sru shrineRegisterUsecase) GetLocnInfoFromPlaceAPI(ctx context.Context, shrrreq *model.ShrineRegisterReq, sac string) (shr *model.Shrine, err error) {

	// PlaceAPIから位置情報(PlaceID、緯度、経度)、及び取得した緯度経度からPlusCodeを取得
	resp, err := placeapi.QueryPlaceAPI(ctx, shrrreq.Name, shrrreq.Address)
	if err != nil {
		return nil, err
	}

	// shrを初期化
	shr = &model.Shrine{}

	// Shrine構造体に値を設定
	shr.Name = shrrreq.Name
	shr.Address = shrrreq.Address
	shr.StdAreaCode = sac
	shr.PlusCode = olc.Encode(shr.Latitude, shr.Longitude, 11)
	shr.Seq = 0
	shr.PlaceID = resp.Results[0].PlaceID
	shr.Latitude = resp.Results[0].Geometry.Location.Lat
	shr.Longitude = resp.Results[0].Geometry.Location.Lng

	fmt.Println(resp.Results[0].Types)

	return shr, nil

}

// 神社テーブル登録
func (sru shrineRegisterUsecase) RegisterShrine(ctx context.Context, shr *model.Shrine) (err error) {

	// 登録するSEQを取得する
	err = sru.sr.GetShrineNextSeq(ctx, shr)
	if err != nil {
		return err
	}

	err = sru.sr.InsertShrine(ctx, shr)
	if err != nil {
		return err
	}

	return nil

}

// 神社詳細情報テーブル登録
func (sru shrineRegisterUsecase) RegisterShrineContents(ctx context.Context, id int, seq int, keyword1 string, keyword2 string, content1 string, content2 string, content3 string, seqHandler int) (err error) {

	// ShrineContents構造体を作成する
	shrc := sru.scr.NewShrineContents(id, seq, keyword1, keyword2, content1, content2, content3)

	// 登録するSEQを取得する
	if seqHandler == 1 {
		err = sru.scr.GetShrineContentsNextSeq(ctx, shrc)
		if err != nil {
			return err
		}
	} else {
		shrc.Seq = 1
	}

	err = sru.scr.InsertShrineContents(ctx, shrc)
	if err != nil {
		return err
	}

	return nil

}

// PlusCodeから神社の登録の有無を判定
func (sru shrineRegisterUsecase) ExistsShrineByPlusCode(ctx context.Context, plusCode string) bool {

	var shrs []*model.Shrine
	var err error

	query := fmt.Sprintf(`SELECT shr.name, shr.address, shr.std_area_code, shr.plus_code, shr.seq, shr.place_id, shr.latitude, shr.longitude, shr.created_at, shr.updated_at
						FROM t_shrines shr
						WHERE shr.plus_code = '%s'`, plusCode)

	shrs, err = sru.sr.GetShrines(ctx, query)
	if err != nil {
		return false
	}

	if len(shrs) == 1 {
		return true
	}

	return false

}

// func SendErrMessageToDiscord(srem model.ShrineRegisterErrMessage) error {

// }
