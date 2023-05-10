package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	Postgres *sql.DB
}

type DatabaseOption struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Dbname  string
	Sslmode string
}

const SSL_DISABLE = "disable"

func ConnectPostgres(option DatabaseOption) (conn DatabaseConnection, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		option.Host, option.Port, option.User, option.Pass, option.Dbname, option.Sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	conn.Postgres = db
	return
}
