package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var conn DatabaseConnection
var err error

func init() {
	option := DatabaseOption{
		Host:    "localhost",
		Port:    "6632",
		User:    "user-search",
		Pass:    "user-pass",
		Dbname:  "search",
		Sslmode: SSL_DISABLE,
	}

	conn, err = ConnectPostgres(option)
}

func TestConnectionPostgresSuccess(t *testing.T) {
	require.Nil(t, err)
	require.NotNil(t, conn.Postgres)
}
