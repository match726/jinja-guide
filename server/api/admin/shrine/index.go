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
	sru := usecase.NewShrineRegisterUsecase(sacp, sp)
	srh := NewShrineRegisterHandler(sru)

	srh.Handler(ctx, w, r)

}

func (srh shrineRegisterHandler) Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	// HTTPリクエストからカスタムヘッダーを取得
	strCustom := r.Header.Get("ShrGuide-Shrines-Authorization")

	// ShrineRegisterReq構造体へ変換
	var shrreq model.ShrineRegisterReq
	err := json.Unmarshal([]byte(strCustom), &shrreq)
	if err != nil {
		logger.Error(ctx, "リクエスト構造体変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	// リクエストされた住所から該当する標準地域コードを取得
	var sac string
	sac, err = srh.sru.GetStdAreaCodeByAddress(ctx, &shrreq)
	if err != nil {
		logger.Error(ctx, "標準地域コード取得失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// PlaceAPIから位置情報(PlaceID、緯度、経度)とPlusCodeを取得
	var shr *model.Shrine
	shr, err = srh.sru.GetLocnInfoFromPlaceAPI(ctx, &shrreq, sac)
	if err != nil {
		logger.Error(ctx, "PlaceAPI取得失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// 神社テーブルへ登録
	err = srh.sru.RegisterShrine(ctx, shr)
	if err != nil {
		logger.Error(ctx, "神社テーブル登録失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
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
