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
