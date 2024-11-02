package api

import (
	"log"

	"github.com/knakk/sparql"
)

func QuerySparql(endpoint string, query string) *sparql.Results {

	repo, err := sparql.NewRepo(endpoint)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	resp, err := repo.Query(query)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return resp

}
