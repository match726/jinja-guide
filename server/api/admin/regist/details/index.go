package api

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/match726/jinja-guide/tree/main/server/logger"
// 	"github.com/match726/jinja-guide/tree/main/server/models"
// 	"github.com/match726/jinja-guide/tree/main/server/trace"
// )

// type ShrineDetailsPostReq struct {
// 	PlusCode        string `json:"plusCode"`
// 	Furigana        string `json:"furigana"`
// 	AltName         string `json:"altName"`
// 	Tags            string `json:"tags"`
// 	FoundedYear     string `json:"foundedYear"`
// 	ObjectOfWorship string `json:"objectOfWorship"`
// 	HasGoshuin      string `json:"hasGoshuin"`
// 	WebsiteURL      string `json:"websiteUrl"`
// 	WikipediaURL    string `json:"wikipediaUrl"`
// }

// func ShrineDetailsRegistHandler(w http.ResponseWriter, r *http.Request) {

// 	switch r.Method {
// 	case http.MethodOptions:
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	case http.MethodPost:
// 		RegisterShrineDetails(w, r)
// 	default:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}

// }

// func RegisterShrineDetails(w http.ResponseWriter, r *http.Request) {

// 	var pg *models.Postgres
// 	var err error

// 	// Contextを生成
// 	ctx := r.Context()
// 	shutdown, err := trace.InitTracerProvider()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer shutdown(ctx)
// 	ctx = trace.GetContextWithTraceID(r.Context(), "RegisterShrineDetails")

// 	// HTTPリクエストからボディを取得
// 	body := make([]byte, r.ContentLength)
// 	r.Body.Read(body)

// 	// ShrineDetailsPostReq構造体へ変換
// 	var shrdpr *ShrineDetailsPostReq
// 	err = json.Unmarshal([]byte(string(body)), &shrdpr)
// 	if err != nil {
// 		fmt.Printf("[Err] <RegisterShrineDetails> Err: パラメータ取得エラー %s\n", err)
// 	}

// 	pg, err = models.NewPool(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	defer pg.ClosePool(ctx)

// 	// 神社の登録があるかをチェック
// 	existsShrine := pg.ExistsShrineByPlusCode(ctx, shrdpr.PlusCode)

// 	if existsShrine {
// 		if len(shrdpr.Furigana) != 0 {
// 			err = pg.InsertShrineContents(ctx, 1, shrdpr.Furigana, shrdpr.PlusCode, 0)
// 			if err != nil {
// 				logger.Error(ctx, "神社詳細情報[振り仮名]登録失敗", "errmsg", err)
// 				w.WriteHeader(http.StatusInternalServerError)
// 			}
// 		}
// 		if len(shrdpr.AltName) != 0 {
// 			err = pg.InsertShrineContents(ctx, 2, shrdpr.AltName, shrdpr.PlusCode, 1)
// 			if err != nil {
// 				logger.Error(ctx, "神社詳細情報[別名称]登録失敗", "errmsg", err)
// 				w.WriteHeader(http.StatusInternalServerError)
// 			}
// 		}
// 		if len(shrdpr.Tags) != 0 {
// 			err = pg.InsertShrineContents(ctx, 4, shrdpr.Tags, shrdpr.PlusCode, 1)
// 			if err != nil {
// 				logger.Error(ctx, "神社詳細情報[関連ワード]登録失敗", "errmsg", err)
// 				w.WriteHeader(http.StatusInternalServerError)
// 			}
// 		}
// 		if len(shrdpr.FoundedYear) != 0 {
// 			err = pg.InsertShrineContents(ctx, 5, shrdpr.FoundedYear, shrdpr.PlusCode, 0)
// 			if err != nil {
// 				logger.Error(ctx, "神社詳細情報[創建年]登録失敗", "errmsg", err)
// 				w.WriteHeader(http.StatusInternalServerError)
// 			}
// 		}
// 		if len(shrdpr.ObjectOfWorship) != 0 {
// 			err = pg.InsertShrineContents(ctx, 6, shrdpr.ObjectOfWorship, shrdpr.PlusCode, 1)
// 			if err != nil {
// 				logger.Error(ctx, "神社詳細情報[御祭神]登録失敗", "errmsg", err)
// 				w.WriteHeader(http.StatusInternalServerError)
// 			}
// 		}
// 		if len(shrdpr.HasGoshuin) != 0 {
// 			err = pg.InsertShrineContents(ctx, 8, shrdpr.HasGoshuin, shrdpr.PlusCode, 0)
// 			if err != nil {
// 				logger.Error(ctx, "神社詳細情報[御朱印]登録失敗", "errmsg", err)
// 				w.WriteHeader(http.StatusInternalServerError)
// 			}
// 		}
// 		if len(shrdpr.WebsiteURL) != 0 {
// 			err = pg.InsertShrineContents(ctx, 9, shrdpr.WebsiteURL, shrdpr.PlusCode, 0)
// 			if err != nil {
// 				logger.Error(ctx, "神社詳細情報[公式サイトURL]登録失敗", "errmsg", err)
// 				w.WriteHeader(http.StatusInternalServerError)
// 			}
// 		}
// 		if len(shrdpr.WikipediaURL) != 0 {
// 			err = pg.InsertShrineContents(ctx, 10, shrdpr.WikipediaURL, shrdpr.PlusCode, 0)
// 			if err != nil {
// 				logger.Error(ctx, "神社詳細情報[WikipediaURL]登録失敗", "errmsg", err)
// 				w.WriteHeader(http.StatusInternalServerError)
// 			}
// 		}
// 	}

// 	writeJsonResp(w, shrdpr)

// }

// func writeJsonResp(w http.ResponseWriter, shrdpr *ShrineDetailsPostReq) {

// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	b, err := json.Marshal(shrdpr)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		s := `{"status":"500 Internal Server Error"}`
// 		if _, err := w.Write([]byte(s)); err != nil {
// 			fmt.Println(err)
// 		}
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	if _, err := w.Write(b); err != nil {
// 		fmt.Println(err)
// 	}

// }
