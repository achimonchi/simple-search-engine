package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var s2 Search
var err2 error

func init() {

	client, _ := ConnectTypesense(SearchOption{
		Host:   "http://localhost:8108",
		APIKey: "ThisIsMasterKey",
	})
	s2.Typesense = client
}

func TestConnectTypesense(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client, err := ConnectTypesense(SearchOption{
			Host:   "http://localhost:8108",
			APIKey: "ThisIsMasterKey",
		})

		require.Nil(t, err)
		require.NotNil(t, client)
	})
	t.Run("error", func(t *testing.T) {
		client, err := ConnectTypesense(SearchOption{
			Host:   "http://localhost:8118",
			APIKey: "ThisIsMasterKey",
		})

		require.NotNil(t, err)
		require.Nil(t, client)
	})
}

func TestMigrateTypesense(t *testing.T) {
	err := s2.MigrateTypesenseUp("../../deploy/data.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	require.Nil(t, err)
}
