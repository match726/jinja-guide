package sparql

import (
	"fmt"

	"github.com/knakk/sparql"
)

func QuerySparql(endpoint string, query string) (*sparql.Results, error) {

	repo, err := sparql.NewRepo(endpoint)
	if err != nil {
		return nil, fmt.Errorf("[SPARQL接続失敗]: %w", err)
	}

	resp, err := repo.Query(query)
	if err != nil {
		return nil, fmt.Errorf("[SPARQLクエリ取得失敗]: %w", err)
	}

	return resp, nil

}
