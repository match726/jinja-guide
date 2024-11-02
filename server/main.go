package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/match726/jinja-guide/tree/main/server/api"
)

var endpoint = "https://jinja-guide-server.vercel.app/"
var origin = "https://jinja-guide.vercel.app/"

func main() {

	//fmt.Println(models.Db)

	db := os.Getenv("POSTGRES_DATABASE")
	fmt.Println(db)
	r := mux.NewRouter()

	//r.HandleFunc("/api/admin/regist", regist).Methods("POST")
	r.HandleFunc("/api/admin/sac", api.Stdareacode).Methods("POST")

	http.ListenAndServe(fmt.Sprintf("%s", endpoint), setHeaders(r))

}

func setHeaders(h http.Handler) http.Handler {

	// CORSを有効にする
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, ShrGuide-User-Authorization, ShrGuide-Shrines-Authorization")
		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)

	})

}
