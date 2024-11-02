package api

import (
	"fmt"
	"net/http"

	"github.com/match726/jinja-guide/tree/main/server/models"
)

func Stdareacode(w http.ResponseWriter, r *http.Request) {
	models.Run().Exit()
	fmt.Fprintf(w, fmt.Sprintf("<p>Hello from Golang!! %s</p>", r))
}
