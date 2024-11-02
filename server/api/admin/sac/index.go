package api

import (
	"fmt"
	"net/http"
)

func Stdareacode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("<p>%s</p>", r))
}
