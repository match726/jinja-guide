package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	logger "github.com/match726/jinja-guide/tree/main/server/infrastructure/log"
	"github.com/match726/jinja-guide/tree/main/server/infrastructure/trace"
	"github.com/match726/jinja-guide/tree/main/server/usecase"
)

type PrefsHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type prefsHandler struct {
	saclu usecase.StdAreaCodeListUsecase
}

// func NewPrefsHandler(saclu usecase.StdAreaCodeListUsecase) PrefsHandler {
// 	return &prefsHandler{saclu: saclu}
// }

func (ph prefsHandler) Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("PrefsHandler")
	fmt.Println("http.Request: %s", r)

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
	ctx = trace.GetContextWithTraceID(r.Context(), "PrefsHandler")

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
