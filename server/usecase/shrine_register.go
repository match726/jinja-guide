package usecase

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	olc "github.com/google/open-location-code/go"
	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/domain/repository"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/discord"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/placeapi"
)

type ShrineRegisterUsecase interface {
	GetAllRegisterShrines(ctx context.Context) (shrrs []*model.ShrineRegister, err error)
	DeleteRegisteredShrine(ctx context.Context, shrq *model.ShrineRegisterReq) (err error)
	GetStdAreaCodeByAddress(ctx context.Context, shrq *model.ShrineRegisterReq) (sac string, err error)
	GetLocnInfoFromPlaceAPI(ctx context.Context, shrq *model.ShrineRegisterReq, sac string) (shr *model.Shrine, caution []string, err error)
	RegisterShrine(ctx context.Context, shr *model.Shrine) (err error)
	RegisterShrineContents(ctx context.Context, id int, seq int, keyword1 string, keyword2 string, content1 string, content2 string, content3 string, seqHandler int) (err error)
	ExistsShrineByPlusCode(ctx context.Context, shrq *model.ShrineRegisterReq) bool
	SendErrMessageToDiscord(procName string, errmsgs []string, shrq *model.ShrineRegisterReq) error
	ConvertSQLErrorMessage(err error) (errmsg string)
}

type shrineRegisterUsecase struct {
	sacr repository.StdAreaCodeRepository
	sr   repository.ShrineRepository
	scr  repository.ShrineContentsRepository
	srr  repository.ShrineRegisterRepository
}

func NewShrineRegisterUsecase(sacr repository.StdAreaCodeRepository, sr repository.ShrineRepository, scr repository.ShrineContentsRepository, srr repository.ShrineRegisterRepository) ShrineRegisterUsecase {
	return &shrineRegisterUsecase{sacr: sacr, sr: sr, scr: scr, srr: srr}
}

// 神社一括登録テーブルから情報取得
func (sru shrineRegisterUsecase) GetAllRegisterShrines(ctx context.Context) (shrrs []*model.ShrineRegister, err error) {

	query := `SELECT rshr.name, rshr.address, rshr.furigana, rshr.alt_name, rshr.tags, rshr.founded_year, rshr.object_of_worship, rshr.shrine_rank, rshr.has_goshuin, rshr.website_url, rshr.wikipedia_url
						FROM m_register_shrine rshr`

	shrrs, err = sru.srr.GetRegisterShrines(ctx, query)
	if err != nil {
		return nil, err
	}

	// for idx, shrr := range shrrs {

	// 	var altNameFields []model.Field
	// 	var tagFields []model.Field
	// 	var foundedYearFields []model.Field
	// 	var objectOfWorshipFields []model.Field
	// 	var shrineRankFields []model.Field

	// 	// 別名称をFieldへ変換
	// 	for _, altName := range shrr.AltNames {
	// 		altNameFields = append(altNameFields, model.Field{Id: idx + 1, Seq: "", Content1: altName, Content2: "", Content3: ""})
	// 	}
	// 	// 関連ワードをFieldへ変換
	// 	for idx, tag := range shrr.Tags {
	// 		tagFields = append(tagFields, model.Field{Id: idx + 1, Seq: "", Content1: tag, Content2: "", Content3: ""})
	// 	}
	// 	// 創建年をFieldへ変換
	// 	if len(shrr.FoundedYear) != 0 {
	// 		foundedYearFields = []model.Field{{Id: 1, Seq: "", Content1: shrr.FoundedYear[0], Content2: shrr.FoundedYear[1], Content3: shrr.FoundedYear[2]}}
	// 	}
	// 	// 御祭神をFieldへ変換
	// 	for idx, objectOfWorships := range shrr.ObjectOfWorships {
	// 		objectOfWorshipFields = append(objectOfWorshipFields, model.Field{Id: idx + 1, Seq: "", Content1: objectOfWorships, Content2: "", Content3: ""})
	// 	}
	// 	// 社格をFieldへ変換
	// 	for idx, shrineRank := range shrr.ShrineRanks {
	// 		if len(shrineRank) == 0 {
	// 			shrineRankFields = append(shrineRankFields, model.Field{Id: idx + 1, Seq: shrineRank[0], Content1: shrineRank[1], Content2: shrineRank[2], Content3: shrineRank[3]})
	// 		}
	// 	}

	// 	shrq := model.ShrineRegisterReq{
	// 		Name:             []model.Field{{Id: 1, Seq: "", Content1: shrr.Name, Content2: "", Content3: ""}},
	// 		Address:          []model.Field{{Id: 1, Seq: "", Content1: shrr.Address, Content2: "", Content3: ""}},
	// 		PlusCode:         []model.Field{{Id: 1, Seq: "", Content1: "", Content2: "", Content3: ""}},
	// 		PlaceID:          []model.Field{{Id: 1, Seq: "", Content1: "", Content2: "", Content3: ""}},
	// 		Furigana:         []model.Field{{Id: 1, Seq: "", Content1: shrr.Furigana, Content2: "", Content3: ""}},
	// 		AltNames:         altNameFields,
	// 		Tags:             tagFields,
	// 		FoundedYear:      foundedYearFields,
	// 		ObjectOfWorships: objectOfWorshipFields,
	// 		ShrineRanks:      shrineRankFields,
	// 		HasGoshuin:       []model.Field{{Id: 1, Seq: "", Content1: shrr.HasGoshuin, Content2: "", Content3: ""}},
	// 		WebsiteURL:       []model.Field{{Id: 1, Seq: "", Content1: shrr.WebsiteURL, Content2: "", Content3: ""}},
	// 		WikipediaURL:     []model.Field{{Id: 1, Seq: "", Content1: shrr.WikipediaURL, Content2: "", Content3: ""}},
	// 	}
	// 	pshrqs = append(pshrqs, &shrq)
	// }

	// fmt.Printf("pshrqs: %+v\n", pshrqs)

	return shrrs, nil

}

// 登録済のレコードを神社一括登録テーブルから削除
func (sru shrineRegisterUsecase) DeleteRegisteredShrine(ctx context.Context, shrq *model.ShrineRegisterReq) (err error) {

	// 登録済のレコードを神社一括登録テーブルから削除
	query := fmt.Sprintf(`DELETE FROM m_register_shrine rshr
						WHERE rshr.name = '%s'
						AND rshr.address = '%s'`, shrq.Name[0].Content1, shrq.Address[0].Content1)

	err = sru.srr.DeleteRegisterShrine(ctx, query)
	if err != nil {
		return err
	}

	return nil

}

// 住所から該当する標準地域コードを取得
func (sru shrineRegisterUsecase) GetStdAreaCodeByAddress(ctx context.Context, shrq *model.ShrineRegisterReq) (sac string, err error) {

	var sacs []*model.StdAreaCode

	// 住所から都道府県を取得
	reg, _ := regexp.Compile(`^東京都|^北海道|^(大阪|京都)府|^\W{2,3}県`)
	pref := reg.FindString(shrq.Address[0].Content1)

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
			if strings.HasPrefix(shrq.Address[0].Content1, keyword) {
				sac = sacs[i].StdAreaCode
				break
			}
		}
	}

	return sac, nil

}

// PlaceAPIから位置情報(PlaceID、緯度、経度)とPlusCodeを取得
// ⇒Shrine構造体の形式で返す
func (sru shrineRegisterUsecase) GetLocnInfoFromPlaceAPI(ctx context.Context, shrq *model.ShrineRegisterReq, sac string) (shr *model.Shrine, caution []string, err error) {

	// PlaceAPIから位置情報(PlaceID、緯度、経度)、及び取得した緯度経度からPlusCodeを取得
	resp, err := placeapi.QueryPlaceAPI(ctx, shrq.Name[0].Content1, shrq.Address[0].Content1)
	if err != nil {
		return nil, []string{}, err
	}

	// Shrine構造体を初期化
	shr = &model.Shrine{}

	// Shrine構造体に値を設定
	shr.Name = shrq.Name[0].Content1
	shr.Address = shrq.Address[0].Content1
	shr.StdAreaCode = sac
	shr.Seq = 0
	shr.PlaceID = resp.Results[0].PlaceID
	shr.Latitude = resp.Results[0].Geometry.Location.Lat
	shr.Longitude = resp.Results[0].Geometry.Location.Lng
	shr.PlusCode = olc.Encode(shr.Latitude, shr.Longitude, 11)

	// ShrineRegisterReq構造体に値を追加
	shrq.PlusCode = []model.Field{{Id: 1, Seq: "", Content1: shr.PlusCode, Content2: "", Content3: ""}}
	shrq.PlaceID = []model.Field{{Id: 1, Seq: "", Content1: shr.PlaceID, Content2: "", Content3: ""}}

	//fmt.Println(resp.Results[0])

	if !slices.Contains(resp.Results[0].Types, "place_of_worship") {
		caution = append(caution, "住所タイプに「place_of_worship」なし")
	}

	if !strings.Contains(resp.Results[0].FormattedAddress, shr.Address) {
		caution = append(caution, "住所不一致（要確認）")
	}

	return shr, caution, nil

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
		shrc.Seq = seq
	}

	err = sru.scr.InsertShrineContents(ctx, shrc)
	if err != nil {
		return err
	}

	return nil

}

// PlusCodeから神社の登録の有無を判定
func (sru shrineRegisterUsecase) ExistsShrineByPlusCode(ctx context.Context, shrq *model.ShrineRegisterReq) bool {

	var shrs []*model.Shrine
	var err error

	// 該当のPlusCodeから神社テーブルを取得
	query := fmt.Sprintf(`SELECT shr.name, shr.address, shr.std_area_code, shr.plus_code, shr.seq, shr.place_id, shr.latitude, shr.longitude, shr.created_at, shr.updated_at
						FROM t_shrines shr
						WHERE shr.plus_code = '%s'`, shrq.PlusCode[0].Content1)

	shrs, err = sru.sr.GetShrines(ctx, query)
	if err != nil {
		return false
	}

	// 登録がある場合、ShrineRegisterReq構造体に値を設定
	if len(shrs) == 1 {
		shrq.Name = []model.Field{{Id: 1, Seq: "", Content1: shrs[0].Name, Content2: "", Content3: ""}}
		shrq.Address = []model.Field{{Id: 1, Seq: "", Content1: shrs[0].Address, Content2: "", Content3: ""}}
		shrq.PlusCode = []model.Field{{Id: 1, Seq: "", Content1: shrs[0].PlusCode, Content2: "", Content3: ""}}
		shrq.PlaceID = []model.Field{{Id: 1, Seq: "", Content1: shrs[0].PlaceID, Content2: "", Content3: ""}}
		return true
	}

	return false

}

// エラー／要確認事象発生時にDiscordへメッセージ送信
func (sru shrineRegisterUsecase) SendErrMessageToDiscord(procName string, errmsgs []string, shrq *model.ShrineRegisterReq) error {

	// エラーメッセージ設定
	content := "[" + procName + "]\n<<エラー概要>>\n"
	for _, errmsg := range errmsgs {
		content = content + "　" + errmsg + "\n"
	}
	content = content + "<<神社情報>>\n"
	if len(shrq.Name) != 0 {
		if len(shrq.Name[0].Content1) != 0 {
			content = content + "　神社名称：" + shrq.Name[0].Content1 + "\n"
		}
	}
	if len(shrq.Address) != 0 {
		if len(shrq.Address[0].Content1) != 0 {
			content = content + "　住所　　：" + shrq.Address[0].Content1 + "\n"
		}
	}
	if len(shrq.PlusCode) != 0 {
		if len(shrq.PlusCode[0].Content1) != 0 {
			content = content + "　PlusCode：" + shrq.PlusCode[0].Content1 + "\n"
		}
	}
	if len(shrq.Name) != 0 && len(shrq.PlaceID) != 0 {
		if len(shrq.Name[0].Content1) != 0 && len(shrq.PlaceID[0].Content1) != 0 {
			content = content + "<<GoogleMapLink>>\nhttps://www.google.com/maps/search/?api=1&query=" + shrq.Name[0].Content1 + "&query_place_id=" + shrq.PlaceID[0].Content1
		}
	}
	err := discord.SendMessage(os.Getenv("DISCORD_ADMIN_WEBHOOK_URL"), os.Getenv("DISCORD_BOT_TOKEN"), content)
	if err != nil {
		return err
	}

	return nil

}

func (sru shrineRegisterUsecase) ConvertSQLErrorMessage(err error) (errmsg string) {

	var sqlcode string

	s := string([]rune(err.Error())[14:])
	// 「SQLSTATE」の開始位置を取得
	posSqlstateSta := strings.Index(s, "SQLSTATE")
	// PostgreSQLエラーコードの開始位置を取得
	posSqlcodeSta := posSqlstateSta + len("SQLSTATE ")

	// 「SQLSTATE」を含む場合エラーコードを取得
	if posSqlstateSta != -1 {
		sqlcode = s[posSqlcodeSta : posSqlcodeSta+5]
	}

	switch sqlcode {
	case "23505":
		errmsg = "既に登録があります"
	default:
		errmsg = "神社テーブル登録失敗"
	}

	return errmsg

}
