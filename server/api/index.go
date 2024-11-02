package api

import (
	"fmt"
	"net/http"
	"os"
)

func Stdareacode(w http.ResponseWriter, r *http.Request) {
	db := os.Getenv("POSTGRES_DATABASE")
	fmt.Println(db)

	fmt.Fprintf(w, fmt.Sprintf("<p>Hello from Golang!! %s</p>", r))
}
