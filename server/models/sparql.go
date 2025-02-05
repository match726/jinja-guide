package models

import (
	"fmt"

	"github.com/knakk/sparql"
)

func QuerySparql(endpoint string, query string) *sparql.Results {

	repo, err := sparql.NewRepo(endpoint)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := repo.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	return resp

}
