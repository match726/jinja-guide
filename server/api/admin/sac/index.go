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

type StdAreaCodeHandler interface {
	GetHandler(ctx context.Context, w http.ResponseWriter, r *http.Request)
	PutHandler(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type stdAreaCodeHandler struct {
	sacu usecase.StdAreaCodeUsecase
}

func NewStdAreaCodeHandler(sacu usecase.StdAreaCodeUsecase) StdAreaCodeHandler {
	return &stdAreaCodeHandler{sacu: sacu}
}

func ExportedHandler(w http.ResponseWriter, r *http.Request) {

	// リクエストメソッド判定
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		break
	case http.MethodPut:
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
	sau := usecase.NewStdAreaCodeUsecase(sacp)
	sach := NewStdAreaCodeHandler(sau)

	// リクエストメソッドでの処理分岐
	switch r.Method {
	case http.MethodGet:
		sach.GetHandler(ctx, w, r)
	case http.MethodPut:
		sach.PutHandler(ctx, w, r)
	}

}

func (sach stdAreaCodeHandler) GetHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	var sacs []*model.StdAreaCode

	// 標準地域コード取得（全件）
	sacs, err := sach.sacu.GetAllStdAreaCodes(ctx)
	if err != nil {
		logger.Error(ctx, "標準地域コード取得失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(sacs)
	if err != nil {
		logger.Error(ctx, "JSON変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		logger.Error(ctx, "Body書込失敗", "errmsg", err)
	}

}

func (sach stdAreaCodeHandler) PutHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	// 標準地域コード最新化
	err := sach.sacu.UpdateStdAreaCode(ctx)
	if err != nil {
		logger.Error(ctx, "標準地域コード最新化失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

}
