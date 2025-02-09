package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/database"
	logger "github.com/match726/jinja-guide/tree/main/server/infrastructure/log"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/persistence"
	tracer "github.com/match726/jinja-guide/tree/main/server/infrastructure/trace"
	"github.com/match726/jinja-guide/tree/main/server/usecase"
)

type ShrineListHandler interface {
	Handler(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type shrineListHandler struct {
	slu usecase.ShrineListUsecase
}

func NewShrineListHandler(slu usecase.ShrineListUsecase) ShrineListHandler {
	return &shrineListHandler{slu: slu}
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
	ctx = tracer.GetContextWithTraceID(r.Context(), "ShrineListHandler")

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
	slp := persistence.NewShrineListPersistence(pg)
	slu := usecase.NewShrineListUsecase(slp)
	slh := NewShrineListHandler(slu)

	slh.Handler(ctx, w, r)

}

func (slh shrineListHandler) Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	// HTTPリクエストからカスタムヘッダーを取得
	strCustom := r.Header.Get("ShrGuide-Shrines-Authorization")

	// ShrineListReq構造体へ変換
	var slreq model.ShrineListReq
	err := json.Unmarshal([]byte(strCustom), &slreq)
	if err != nil {
		logger.Error(ctx, "リクエスト構造体変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// リクエストパラメータチェック
	if len(slreq.Tag) != 0 {
		logger.Error(ctx, "リクエストパラメータ不正検知", "errmsg", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// リクエストパラメータをデコード
	slreq.Tag, _ = url.QueryUnescape(slreq.Tag)

	// 神社一覧（タグ単位）を取得
	var slrsps []*model.ShrineListResp
	slrsps, err = slh.slu.GetShrineListByTag(ctx, slreq.Tag)
	if err != nil {
		logger.Error(ctx, "神社一覧（タグ単位）取得失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println(slrsps)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(slrsps)
	if err != nil {
		logger.Error(ctx, "JSON変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		logger.Error(ctx, "Body書込失敗", "errmsg", err)
	}

}
