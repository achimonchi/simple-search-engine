package db

import (
	"database/sql"
	"fmt"
	"log"

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

var createProducts = `
CREATE TABLE "products" (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL,
    description text NOT NULL,
    price float NOT NULL,
    stock int NOT NULL,
    created_at timestamptz DEFAULT NOW()
);
`

func (d DatabaseConnection) MigratePostgres() {
	log.Println("running db migration")
	_, err := d.Postgres.Exec(createProducts)
	if err != nil {
		panic(err)
	}
	log.Println("migration done")
}
