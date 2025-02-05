package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/match726/jinja-guide/tree/main/server/infrastructure/database"
	logger "github.com/match726/jinja-guide/tree/main/server/infrastructure/log"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/persistence"
	tracer "github.com/match726/jinja-guide/tree/main/server/infrastructure/trace"
	"github.com/match726/jinja-guide/tree/main/server/usecase"
)

type PrefsHandler interface {
	Handler(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type prefsHandler struct {
	saclu usecase.StdAreaCodeListUsecase
}

func NewPrefsHandler(saclu usecase.StdAreaCodeListUsecase) PrefsHandler {
	return &prefsHandler{saclu: saclu}
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
	saclp := persistence.NewStdAreaCodeListPersistence(pg)
	saclu := usecase.NewStdAreaCodeListUsecase(saclp)
	ph := NewPrefsHandler(saclu)

	ph.Handler(ctx, w, r)

}

func (ph prefsHandler) Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	// 登録されている神社を元に都道府県の標準地域コード（紐付き）を取得する
	sacrrs, err := ph.saclu.GetAllStdAreaCodeRelationshipList(ctx)
	if err != nil {
		logger.Error(ctx, "標準地域コード（紐付き）取得失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(sacrrs)
	if err != nil {
		logger.Error(ctx, "JSON変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		logger.Error(ctx, "Body書込失敗", "errmsg", err)
	}

}
