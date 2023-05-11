package search

import (
	"errors"

	"github.com/meilisearch/meilisearch-go"
)

type SearchOption struct {
	Host   string
	APIKey string
}

type Search struct {
	Meilisearch *meilisearch.Client
}

func ConnectMeili(option SearchOption) (s Search, err error) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   option.Host, // like http://localhost:7700
		APIKey: option.APIKey,
	})

	if !client.IsHealthy() {
		err = errors.New("meilisearch not healthy")
		return
	}
	s.Meilisearch = client
	return
}
