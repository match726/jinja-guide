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

type ShrineDetailHandler interface {
	Handler(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type shrineDetailHandler struct {
	sdu usecase.ShrineDetailUsecase
}

func NewShrineDetailHandler(sdu usecase.ShrineDetailUsecase) ShrineDetailHandler {
	return &shrineDetailHandler{sdu: sdu}
}

func ExportedHandler(w http.ResponseWriter, r *http.Request) {

	// リクエストメソッド判定
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
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
	ctx = tracer.GetContextWithTraceID(r.Context(), "PrefsHandler")

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
	shrp := persistence.NewShrinePersistence(pg)
	shrcp := persistence.NewShrineContentsPersistence(pg)
	sdu := usecase.NewShrineDetailUsecase(shrp, shrcp)
	sdh := NewShrineDetailHandler(sdu)

	sdh.Handler(ctx, w, r)

}

func (sdh shrineDetailHandler) Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	// HTTPリクエストからカスタムヘッダーを取得
	strCustom := r.Header.Get("ShrGuide-Shrines-Authorization")

	// ShrineDetailsReq構造体へ変換
	var shrdreq model.ShrineDetailsReq
	err := json.Unmarshal([]byte(strCustom), &shrdreq)
	if err != nil {
		logger.Error(ctx, "リクエスト構造体変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// 神社詳細画面のレスポンス用データを取得
	var shrdrsp *model.ShrineDetailsResp
	shrdrsp, err = sdh.sdu.GetShrineDetailByPlusCode(ctx, shrdreq.PlusCode)
	if err != nil {
		logger.Error(ctx, "神社詳細情報取得失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(shrdrsp)
	if err != nil {
		logger.Error(ctx, "JSON変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		logger.Error(ctx, "Body書込失敗", "errmsg", err)
	}

}
