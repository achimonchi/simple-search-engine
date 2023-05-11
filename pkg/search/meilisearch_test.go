package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var s Search
var err error

func init() {

}

func TestConnectMeilisearch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s, err = ConnectMeili(SearchOption{
			Host:   "http://localhost:7700",
			APIKey: "ThisIsMasterKey",
		})

		require.Nil(t, err)
		require.NotNil(t, s.Meilisearch)

	})
	t.Run("error", func(t *testing.T) {
		s, err = ConnectMeili(SearchOption{
			Host:   "http://localhost:770",
			APIKey: "ThisIsMasterKey",
		})

		require.NotNil(t, err)
		require.Equal(t, err.Error(), "meilisearch not healthy")

	})
}
