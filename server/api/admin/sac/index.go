package api

import (
	"fmt"
	"net/http"
)

func StdAreaCodeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		FetchStdAreaCodes(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func FetchStdAreaCodes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("標準地域コード取得")
}
