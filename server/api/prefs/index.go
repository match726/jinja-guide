package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/match726/jinja-guide/tree/main/server/infrastructure/database"
	logger "github.com/match726/jinja-guide/tree/main/server/infrastructure/log"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/persistence"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/trace"
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

	// Contextを生成
	ctx := r.Context()
	shutdown, err := trace.InitTracerProvider()
	if err != nil {
		logger.Error(ctx, "トレーサープロバイダー作成失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer shutdown(ctx)
	ctx = trace.GetContextWithTraceID(r.Context(), "PrefsHandler")

	var pg *database.Postgres

	pg, err = database.NewPool(ctx)
	if err != nil {
		logger.Error(ctx, "コネクションプール作成失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer pg.ClosePool(ctx)

	saclp := persistence.NewStdAreaCodeListPersistence(pg)
	saclu := usecase.NewStdAreaCodeListUsecase(saclp)
	ph := NewPrefsHandler(saclu)

	ph.Handler(ctx, w, r)

}

func (ph prefsHandler) Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	fmt.Println("Handler")

	// 登録されている神社を元に都道府県の標準地域コード（紐付き）を取得する
	sacrrs, err := ph.saclu.GetAllStdAreaCodeRelationshipList(ctx)
	if err != nil {
		logger.Error(ctx, "標準地域コード（紐付き）取得失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println(sacrrs)

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
