package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

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
	srp := persistence.NewShrineRegisterPersistence(pg)
	sru := usecase.NewShrineRegisterUsecase(sacp, sp, scp, srp)
	srh := NewShrineRegisterHandler(sru)

	srh.Handler(ctx, w, r)

}

func (srh shrineRegisterHandler) Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	// HTTPリクエストからボディを取得
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	// ShrineRegisterReq構造体へ変換
	shrq := &model.ShrineRegisterReq{}
	err := json.Unmarshal([]byte(string(body)), shrq)
	if err != nil {
		logger.Error(ctx, "リクエスト構造体変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 神社の登録があるかをチェック
	existsShrine := srh.sru.ExistsShrineByPlusCode(ctx, shrq)
	if !existsShrine {
		logger.Error(ctx, "対象神社検索失敗", "errmsg", err)
		err = srh.sru.SendErrMessageToDiscord("神社詳細情報登録", []string{"対象神社検索失敗"}, shrq)
		if err != nil {
			logger.Error(ctx, "Discord連携失敗", "errmsg", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(shrq.Furigana) != 0 {
		if len(shrq.Furigana[0].Content1) != 0 {
			err = srh.sru.RegisterShrineContents(ctx, 1, 1, shrq.PlusCode[0].Content1, "", shrq.Furigana[0].Content1, "", "", 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[振り仮名]登録失敗", "errmsg", err)
				err = srh.sru.SendErrMessageToDiscord("神社詳細情報登録", []string{"神社詳細情報[振り仮名]登録失敗"}, shrq)
				if err != nil {
					logger.Error(ctx, "Discord連携失敗", "errmsg", err)
				}
			}
		}
	}
	if len(shrq.AltNames) != 0 {
		for _, altName := range shrq.AltNames {
			if len(altName.Content1) != 0 {
				err = srh.sru.RegisterShrineContents(ctx, 2, 1, shrq.PlusCode[0].Content1, "", altName.Content1, "", "", 1)
				if err != nil {
					logger.Error(ctx, "神社詳細情報[別名称]登録失敗", "errmsg", err)
					err = srh.sru.SendErrMessageToDiscord("神社詳細情報登録", []string{"神社詳細情報[別名称]登録失敗"}, shrq)
					if err != nil {
						logger.Error(ctx, "Discord連携失敗", "errmsg", err)
					}
				}
			}
		}
	}
	if len(shrq.Tags) != 0 {
		for _, tag := range shrq.Tags {
			if len(tag.Content1) != 0 {
				err = srh.sru.RegisterShrineContents(ctx, 4, 1, shrq.PlusCode[0].Content1, "", tag.Content1, "", "", 1)
				if err != nil {
					logger.Error(ctx, "神社詳細情報[関連ワード]登録失敗", "errmsg", err)
					err = srh.sru.SendErrMessageToDiscord("神社詳細情報登録", []string{"神社詳細情報[関連ワード]登録失敗"}, shrq)
					if err != nil {
						logger.Error(ctx, "Discord連携失敗", "errmsg", err)
					}
				}
			}
		}
	}
	if len(shrq.FoundedYear) != 0 {
		if len(shrq.FoundedYear[0].Content1) != 0 {
			err = srh.sru.RegisterShrineContents(ctx, 5, 1, shrq.PlusCode[0].Content1, "", shrq.FoundedYear[0].Content1, shrq.FoundedYear[0].Content2, "", 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[創建年]登録失敗", "errmsg", err)
				err = srh.sru.SendErrMessageToDiscord("神社詳細情報登録", []string{"神社詳細情報[創建年]登録失敗"}, shrq)
				if err != nil {
					logger.Error(ctx, "Discord連携失敗", "errmsg", err)
				}
			}
		}
	}
	if len(shrq.ObjectOfWorships) != 0 {
		for _, objectOfWorship := range shrq.ObjectOfWorships {
			if len(objectOfWorship.Content1) != 0 {
				err = srh.sru.RegisterShrineContents(ctx, 6, 1, shrq.PlusCode[0].Content1, "", objectOfWorship.Content1, "", "", 1)
				if err != nil {
					logger.Error(ctx, "神社詳細情報[御祭神]登録失敗", "errmsg", err)
					err = srh.sru.SendErrMessageToDiscord("神社詳細情報登録", []string{"神社詳細情報[御祭神]登録失敗"}, shrq)
					if err != nil {
						logger.Error(ctx, "Discord連携失敗", "errmsg", err)
					}
				}
			}
		}
	}
	if len(shrq.ShrineRanks) != 0 {
		for _, shrineRank := range shrq.ShrineRanks {
			if len(shrineRank.Seq) != 0 && len(shrineRank.Content1) != 0 {
				// Seqをint型に変換
				seq, _ := strconv.Atoi(shrineRank.Seq)
				err = srh.sru.RegisterShrineContents(ctx, 7, seq, shrq.PlusCode[0].Content1, "", shrineRank.Content1, shrineRank.Content2, "", 0)
				if err != nil {
					logger.Error(ctx, "神社詳細情報[社格]登録失敗", "errmsg", err)
					err = srh.sru.SendErrMessageToDiscord("神社詳細情報登録", []string{"神社詳細情報[社格]登録失敗"}, shrq)
					if err != nil {
						logger.Error(ctx, "Discord連携失敗", "errmsg", err)
					}
				}
			}
		}
	}
	if len(shrq.HasGoshuin) != 0 {
		if len(shrq.HasGoshuin[0].Content1) != 0 {
			err = srh.sru.RegisterShrineContents(ctx, 8, 1, shrq.PlusCode[0].Content1, "", shrq.HasGoshuin[0].Content1, "", "", 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[御朱印]登録失敗", "errmsg", err)
				err = srh.sru.SendErrMessageToDiscord("神社詳細情報登録", []string{"神社詳細情報[御朱印]登録失敗"}, shrq)
				if err != nil {
					logger.Error(ctx, "Discord連携失敗", "errmsg", err)
				}
			}
		}
	}
	if len(shrq.WebsiteURL) != 0 {
		if len(shrq.WebsiteURL[0].Content1) != 0 {
			err = srh.sru.RegisterShrineContents(ctx, 9, 1, shrq.PlusCode[0].Content1, "", shrq.WebsiteURL[0].Content1, "", "", 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[公式サイトURL]登録失敗", "errmsg", err)
				err = srh.sru.SendErrMessageToDiscord("神社詳細情報登録", []string{"神社詳細情報[公式サイトURL]登録失敗"}, shrq)
				if err != nil {
					logger.Error(ctx, "Discord連携失敗", "errmsg", err)
				}
			}
		}
	}
	if len(shrq.WikipediaURL) != 0 {
		if len(shrq.WikipediaURL[0].Content1) != 0 {
			err = srh.sru.RegisterShrineContents(ctx, 10, 1, shrq.PlusCode[0].Content1, "", shrq.WikipediaURL[0].Content1, "", "", 0)
			if err != nil {
				logger.Error(ctx, "神社詳細情報[WikipediaURL]登録失敗", "errmsg", err)
				err = srh.sru.SendErrMessageToDiscord("神社詳細情報登録", []string{"神社詳細情報[WikipediaURL]登録失敗"}, shrq)
				if err != nil {
					logger.Error(ctx, "Discord連携失敗", "errmsg", err)
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(shrq)

	if err != nil {
		logger.Error(ctx, "JSON変換失敗", "errmsg", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		logger.Error(ctx, "Body書込失敗", "errmsg", err)
	}

}
