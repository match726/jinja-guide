package api

import (
	"fmt"
	"net/http"
)

func Stdareacode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Golang!!</h1>")
}
