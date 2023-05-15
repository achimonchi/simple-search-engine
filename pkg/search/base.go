package search

import (
	"github.com/meilisearch/meilisearch-go"
	"github.com/typesense/typesense-go/typesense"
)

type SearchOption struct {
	Host   string
	APIKey string
}

type Search struct {
	Meilisearch *meilisearch.Client
	Typesense   *typesense.Client
}

func NewSearchEngine() Search {
	return Search{}
}

func (s Search) SetMeilisearch(client *meilisearch.Client) Search {
	s.Meilisearch = client
	return s
}

func (s Search) SetTypesense(client *typesense.Client) Search {
	s.Typesense = client
	return s
}
