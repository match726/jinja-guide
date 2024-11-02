package api

import (
	"fmt"
	"net/http"
)

func Stdareacode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("<p>Hello from Golang!! %s</p>", r))
}
