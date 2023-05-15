package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var s Search
var err error

func init() {
	client, _ := ConnectMeili(SearchOption{
		Host:   "http://localhost:7700",
		APIKey: "ThisIsMasterKey",
	})

	s.Meilisearch = client
}

func TestConnectMeilisearch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s, err := ConnectMeili(SearchOption{
			Host:   "http://localhost:7700",
			APIKey: "ThisIsMasterKey",
		})

		require.Nil(t, err)
		require.NotNil(t, s)

	})
	t.Run("error", func(t *testing.T) {
		_, err := ConnectMeili(SearchOption{
			Host:   "http://localhost:770",
			APIKey: "ThisIsMasterKey",
		})

		require.NotNil(t, err)
		require.Equal(t, err.Error(), "meilisearch not healthy")

	})
}

func TestInsertData(t *testing.T) {
	documents := []map[string]interface{}{
		{
			"title":  "",
			"genres": "comedy",
			"id":     22123,
		},
	}
	s.Meilisearch.Index("products").UpdateDocuments(documents)
}
