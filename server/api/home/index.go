package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/match726/jinja-guide/tree/main/server/infrastructure/database"
	logger "github.com/match726/jinja-guide/tree/main/server/infrastructure/log"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/persistence"
	tracer "github.com/match726/jinja-guide/tree/main/server/infrastructure/trace"
	"github.com/match726/jinja-guide/tree/main/server/usecase"
)

type HomeHandler interface {
	Handler(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type homeHandler struct {
	hcu usecase.HomeContentsUsecase
}

func NewHomeHandler(hcu usecase.HomeContentsUsecase) HomeHandler {
	return &homeHandler{hcu: hcu}
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
	hcu := usecase.NewHomeContentsUsecase(shrp, shrcp)
	hh := NewHomeHandler(hcu)

	hh.Handler(ctx, w, r)

}

func (hh homeHandler) Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	// ホーム画面レスポンス取得
	hcrsp, err := hh.hcu.GetHomeContents(ctx)
	if err != nil {
		logger.Error(ctx, "ホーム画面レスポンス取得失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println(hcrsp)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(hcrsp)
	if err != nil {
		logger.Error(ctx, "JSON変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		logger.Error(ctx, "Body書込失敗", "errmsg", err)
	}

}
