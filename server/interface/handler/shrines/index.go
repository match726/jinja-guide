package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
	logger "github.com/match726/jinja-guide/tree/main/server/infrastructure/log"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/trace"
	"github.com/match726/jinja-guide/tree/main/server/usecase"
)

type ShrineListHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type shrineListHandler struct {
	slu usecase.ShrineListUsecase
}

func NewShrineListHandler(slu usecase.ShrineListUsecase) ShrineListHandler {
	return &shrineListHandler{slu: slu}
}

func (slh shrineListHandler) Handler(w http.ResponseWriter, r *http.Request) {

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
	}
	defer shutdown(ctx)
	ctx = trace.GetContextWithTraceID(r.Context(), "ShrineListHandler")

	// HTTPリクエストからカスタムヘッダーを取得
	strCustom := r.Header.Get("ShrGuide-Shrines-Authorization")

	// ShrineListReq構造体へ変換
	var slreq model.ShrineListReq
	err = json.Unmarshal([]byte(strCustom), &slreq)
	if err != nil {
		logger.Error(ctx, "リクエスト構造体変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var slrsps []*model.ShrineListResp
	if len(slreq.Kinds) != 0 && len(slreq.StdAreaCode) != 0 {
		slrsps, err = slh.slu.GetShrineListByStdAreaCode(ctx, slreq.Kinds, slreq.StdAreaCode)
		if err != nil {
			logger.Error(ctx, "神社一覧（都道府県単位）取得失敗", "errmsg", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else if len(slreq.Tag) != 0 {
		slrsps, err = slh.slu.GetShrineListByTag(ctx, slreq.Tag)
		if err != nil {
			logger.Error(ctx, "神社一覧（タグ単位）取得失敗", "errmsg", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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
