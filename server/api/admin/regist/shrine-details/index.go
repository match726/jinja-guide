package api

import (
	"context"
	"encoding/json"
	"net/http"

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
	sru := usecase.NewShrineRegisterUsecase(sacp, sp, scp)
	srh := NewShrineRegisterHandler(sru)

	srh.Handler(ctx, w, r)

}

func (srh shrineRegisterHandler) Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	// HTTPリクエストからボディを取得
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	// ShrineContentsRegisterReq構造体へ変換
	var shrcreq model.ShrineContentsRegisterReq
	err := json.Unmarshal([]byte(string(body)), &shrcreq)
	if err != nil {
		logger.Error(ctx, "リクエスト構造体変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// 神社の登録があるかをチェック
	existsShrine := srh.sru.ExistsShrineByPlusCode(ctx, shrcreq.PlusCode)
	if !existsShrine {
		logger.Error(ctx, "対象神社検索不可", "errmsg", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(shrcreq.Furigana) != 0 {
		err = srh.sru.RegisterShrineContents(ctx, 1, 1, shrcreq.PlusCode, "", shrcreq.Furigana, "", "", 0)
		if err != nil {
			logger.Error(ctx, "神社詳細情報[振り仮名]登録失敗", "errmsg", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	if len(shrcreq.AltName) != 0 {
		err = srh.sru.RegisterShrineContents(ctx, 2, 1, shrcreq.PlusCode, "", shrcreq.AltName, "", "", 1)
		if err != nil {
			logger.Error(ctx, "神社詳細情報[別名称]登録失敗", "errmsg", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	if len(shrcreq.Tags) != 0 {
		err = srh.sru.RegisterShrineContents(ctx, 4, 1, shrcreq.PlusCode, "", shrcreq.Tags, "", "", 1)
		if err != nil {
			logger.Error(ctx, "神社詳細情報[関連ワード]登録失敗", "errmsg", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	if len(shrcreq.FoundedYear) != 0 {
		err = srh.sru.RegisterShrineContents(ctx, 5, 1, shrcreq.PlusCode, "", shrcreq.FoundedYear, "", "", 0)
		if err != nil {
			logger.Error(ctx, "神社詳細情報[創建年]登録失敗", "errmsg", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	if len(shrcreq.ObjectOfWorship) != 0 {
		err = srh.sru.RegisterShrineContents(ctx, 6, 1, shrcreq.PlusCode, "", shrcreq.ObjectOfWorship, "", "", 1)
		if err != nil {
			logger.Error(ctx, "神社詳細情報[御祭神]登録失敗", "errmsg", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	if len(shrcreq.HasGoshuin) != 0 {
		err = srh.sru.RegisterShrineContents(ctx, 8, 1, shrcreq.PlusCode, "", shrcreq.HasGoshuin, "", "", 0)
		if err != nil {
			logger.Error(ctx, "神社詳細情報[御朱印]登録失敗", "errmsg", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	if len(shrcreq.WebsiteURL) != 0 {
		err = srh.sru.RegisterShrineContents(ctx, 9, 1, shrcreq.PlusCode, "", shrcreq.WebsiteURL, "", "", 0)
		if err != nil {
			logger.Error(ctx, "神社詳細情報[公式サイトURL]登録失敗", "errmsg", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	if len(shrcreq.WikipediaURL) != 0 {
		err = srh.sru.RegisterShrineContents(ctx, 10, 1, shrcreq.PlusCode, "", shrcreq.WikipediaURL, "", "", 0)
		if err != nil {
			logger.Error(ctx, "神社詳細情報[WikipediaURL]登録失敗", "errmsg", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(shrcreq)

	if err != nil {
		logger.Error(ctx, "JSON変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		logger.Error(ctx, "Body書込失敗", "errmsg", err)
	}

}
