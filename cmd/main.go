package main

import (
	"flag"
	"meilisearch/pkg/db"
	"meilisearch/pkg/search"
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

	searchClient, err := search.ConnectMeili(search.SearchOption{
		Host:   "http://localhost:7700",
		APIKey: "ThisIsMasterKey",
	})

	if err != nil {
		panic(err)
	}

	migrateUp := flag.Bool("migrate-up", false, "setup migration for golang")
	migrateDown := flag.Bool("migrate-down", false, "setup migration for golang")

	flag.Parse()

	// setup migration up
	// will create several tables
	if *migrateUp {
		dbConn.MigratePostgres()
	}

	// setup migration down
	// will remove several tables
	if *migrateDown {
		dbConn.RemoveTablePostgres()
	}

	// if no migrate setup, will running API server
	if !*migrateUp && !*migrateDown {
		myAPI := api.NewAPI().
			SetDatabase(dbConn).
			SetPort(":8888").
			SetMaxProcess(1).
			SetSearchClient(searchClient)

		myAPI.GenerateRoute()
	}

}
