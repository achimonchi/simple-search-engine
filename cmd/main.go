package main

import (
	"meilisearch/pkg/db"
	"meilisearch/usecase/api"
)

func main() {
	dbConn, err := db.ConnectPostgres(db.DatabaseOption{
		Host:    "localhost",
		Port:    "6632",
		User:    "user-search",
		Pass:    "user-pass",
		Dbname:  "search",
		Sslmode: db.SSL_DISABLE,
	})

	if err != nil {
		panic(err)
	}

	myAPI := api.NewAPI().SetDatabase(dbConn).SetPort(":8888").SetMaxProcess(1)

	myAPI.GenerateRoute()
}
