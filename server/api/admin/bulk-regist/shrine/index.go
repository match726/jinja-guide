package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/database"
	logger "github.com/match726/jinja-guide/tree/main/server/infrastructure/log"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/persistence"
	tracer "github.com/match726/jinja-guide/tree/main/server/infrastructure/trace"
	"github.com/match726/jinja-guide/tree/main/server/usecase"
)

type ShrineRegisterHandler interface {
	Handler(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type shrineRegisterHandler struct {
	sru usecase.ShrineRegisterUsecase
}

func NewShrineRegisterHandler(sru usecase.ShrineRegisterUsecase) ShrineRegisterHandler {
	return &shrineRegisterHandler{sru: sru}
}

func ExportedHandler(w http.ResponseWriter, r *http.Request) {

	// リクエストメソッド判定
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodPost:
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Context生成、TraceID、SpanID取得
	ctx := r.Context()
	shutdown, err := tracer.InitTracerProvider()
	if err != nil {
		logger.Error(ctx, "トレーサープロバイダー作成失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer shutdown(ctx)
	ctx = tracer.GetContextWithTraceID(r.Context(), "ShrineRegisterHandler")

	// コネクションプール作成
	var pg *database.Postgres
	pg, err = database.NewPool(ctx)
	if err != nil {
		logger.Error(ctx, "コネクションプール作成失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer pg.ClosePool(ctx)

	// 依存性注入（DI）
	sacp := persistence.NewStdAreaCodePersistence(pg)
	sp := persistence.NewShrinePersistence(pg)
	scp := persistence.NewShrineContentsPersistence(pg)
	srp := persistence.NewShrineRegisterPersistence(pg)
	sru := usecase.NewShrineRegisterUsecase(sacp, sp, scp, srp)
	srh := NewShrineRegisterHandler(sru)

	srh.Handler(ctx, w, r)

}

func (srh shrineRegisterHandler) Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	// 神社一括登録テーブルを取得
	var shrreqs []*model.ShrineRegisterReq
	shrreqs, err := srh.sru.GetAllRegisterShrines(ctx)
	if err != nil {
		logger.Error(ctx, "神社一括登録テーブル取得失敗", "errmsg", err)
		err = srh.sru.SendErrMessageToDiscord([]string{"神社一括登録テーブル取得失敗"}, nil, nil)
		if err != nil {
			logger.Error(ctx, "Discord連携失敗", "errmsg", err)
		}
		w.WriteHeader(http.StatusBadRequest)
	}

	var sac string
	var shr *model.Shrine
	var caution []string

	for _, shrreq := range shrreqs {

		// リクエストされた住所から該当する標準地域コードを取得
		sac, err = srh.sru.GetStdAreaCodeByAddress(ctx, shrreq)
		if err != nil {
			logger.Error(ctx, "標準地域コード取得失敗", "errmsg", err)
			err = srh.sru.SendErrMessageToDiscord([]string{"標準地域コード取得失敗"}, shrreq, nil)
			if err != nil {
				logger.Error(ctx, "Discord連携失敗", "errmsg", err)
			}
			w.WriteHeader(http.StatusInternalServerError)
		}

		// PlaceAPIから位置情報(PlaceID、緯度、経度)とPlusCodeを取得
		shr, caution, err = srh.sru.GetLocnInfoFromPlaceAPI(ctx, shrreq, sac)
		if err != nil {
			logger.Error(ctx, "PlaceAPI取得失敗", "errmsg", err)
			err = srh.sru.SendErrMessageToDiscord([]string{"PlaceAPI取得失敗"}, shrreq, shr)
			if err != nil {
				logger.Error(ctx, "Discord連携失敗", "errmsg", err)
			}
			w.WriteHeader(http.StatusInternalServerError)
		}

		if len(caution) != 0 {
			err = srh.sru.SendErrMessageToDiscord(caution, shrreq, shr)
			if err != nil {
				logger.Error(ctx, "Discord連携失敗", "errmsg", err)
			}
		}

		if len(shr.PlusCode) == 0 {
			err = srh.sru.SendErrMessageToDiscord([]string{"PlusCode取得失敗"}, shrreq, shr)
			if err != nil {
				logger.Error(ctx, "Discord連携失敗", "errmsg", err)
			}
			w.WriteHeader(http.StatusInternalServerError)
		}

		// 神社テーブルへ登録
		err = srh.sru.RegisterShrine(ctx, shr)
		if err != nil {
			logger.Error(ctx, "神社テーブル登録失敗", "errmsg", err)
			err = srh.sru.SendErrMessageToDiscord([]string{srh.sru.ConvertSQLErrorMessage(err)}, shrreq, shr)
			if err != nil {
				logger.Error(ctx, "Discord連携失敗", "errmsg", err)
			}
			w.WriteHeader(http.StatusInternalServerError)
		}

		// 神社詳細テーブルへ登録
		if len(shrreq.Furigana) != 0 {
			err = srh.sru.RegisterShrineContents(ctx, 1, 1, shr.PlusCode, "", shrreq.Furigana, "", "", 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[振り仮名]登録失敗", "errmsg", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		if len(shrreq.AltName) != 0 {
			for i := 0; i < len(shrreq.AltName); i++ {
				if len(strings.TrimSpace(shrreq.AltName[i])) != 0 {
					err = srh.sru.RegisterShrineContents(ctx, 2, 1, shr.PlusCode, "", shrreq.AltName[i], "", "", 1)
					if err != nil {
						logger.Error(ctx, "神社詳細情報[別名称]登録失敗", "errmsg", err)
						w.WriteHeader(http.StatusInternalServerError)
					}
				}
			}
		}
		if len(shrreq.Tags) != 0 {
			for i := 0; i < len(shrreq.Tags); i++ {
				err = srh.sru.RegisterShrineContents(ctx, 4, 1, shr.PlusCode, "", shrreq.Tags[i], "", "", 1)
				if err != nil {
					logger.Error(ctx, "神社詳細情報[関連ワード]登録失敗", "errmsg", err)
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}
		if len(shrreq.FoundedYear) != 0 {
			err = srh.sru.RegisterShrineContents(ctx, 5, 1, shr.PlusCode, "", shrreq.FoundedYear, "", "", 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[創建年]登録失敗", "errmsg", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		if len(shrreq.ObjectOfWorship) != 0 {
			for i := 0; i < len(shrreq.ObjectOfWorship); i++ {
				err = srh.sru.RegisterShrineContents(ctx, 6, 1, shr.PlusCode, "", shrreq.ObjectOfWorship[i], "", "", 1)
				if err != nil {
					logger.Error(ctx, "神社詳細情報[御祭神]登録失敗", "errmsg", err)
					w.WriteHeader(http.StatusInternalServerError)
				}
			}
		}
		if len(shrreq.HasGoshuin) != 0 {
			err = srh.sru.RegisterShrineContents(ctx, 8, 1, shr.PlusCode, "", shrreq.HasGoshuin, "", "", 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[御朱印]登録失敗", "errmsg", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		if len(shrreq.WebsiteURL) != 0 {
			err = srh.sru.RegisterShrineContents(ctx, 9, 1, shr.PlusCode, "", shrreq.WebsiteURL, "", "", 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[公式サイトURL]登録失敗", "errmsg", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		if len(shrreq.WikipediaURL) != 0 {
			err = srh.sru.RegisterShrineContents(ctx, 10, 1, shr.PlusCode, "", shrreq.WikipediaURL, "", "", 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[WikipediaURL]登録失敗", "errmsg", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

		// 登録完了レコードを削除
		err = srh.sru.DeleteRegisteredShrine(ctx, shrreq)
		if err != nil {
			logger.Error(ctx, "神社一括登録テーブル削除失敗", "errmsg", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(shr)
	if err != nil {
		logger.Error(ctx, "JSON変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		logger.Error(ctx, "Body書込失敗", "errmsg", err)
	}

}
