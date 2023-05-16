package main

import (
	"flag"
	"log"
	"meilisearch/pkg/db"
	"meilisearch/pkg/search"
	"meilisearch/usecase/api"
	"os"
)

func main() {
	dbConn, err := db.ConnectPostgres(db.DatabaseOption{
		Host:    getConfig("DB_HOST", "localhost"),
		Port:    getConfig("DB_PORT", "6632"),
		User:    getConfig("DB_USER", "user-search"),
		Pass:    getConfig("DB_PASS", "user-pass"),
		Dbname:  getConfig("DB_NAME", "search"),
		Sslmode: db.SSL_DISABLE,
	})

	if err != nil {
		panic(err)
	}

	searchClient := search.NewSearchEngine()

	// setup meilisearch
	meiliClient, err := search.ConnectMeili(search.SearchOption{
		Host:   getConfig("MEILI_HOST", "http://localhost:7700"),
		APIKey: getConfig("MEILI_APIKEY", "ThisIsMasterKey"),
	})

	if err != nil {
		panic(err)
	}

	// setup typesense search
	typesenseClient, err := search.ConnectTypesense(search.SearchOption{
		Host:   getConfig("TYPESENSE_HOST", "http://localhost:8108"),
		APIKey: getConfig("TYPESENSE_APIKEY", "ThisIsMasterKey"),
	})
	if err != nil {
		panic(err)
	}

	searchClient = searchClient.
		SetMeilisearch(meiliClient).
		SetTypesense(typesenseClient)

	migrate := flag.String("migrate", "", "setup migration for golang. you can use `up` or `down`")
	migrateSearch := flag.String("migrate-search", "", "setup migration for search engine")

	flag.Parse()

	// setup migration up
	// will create several tables
	if *migrate == "up" {
		dbConn.MigratePostgres()
	} else if *migrate == "down" {
		dbConn.RemoveTablePostgres()
	}

	if *migrateSearch == "up" {
		err := searchClient.MigrateUp("deploy/data.json")
		if err != nil {
			log.Println("error when try to migrate search data with error :", err)
		}
		err = searchClient.MigrateTypesenseUp("deploy/data.json")
		if err != nil {
			log.Println("error when try to migrate search data with error :", err)
		}
	}

	// if no migrate setup, will running API server
	if *migrate == "" && *migrateSearch == "" {
		myAPI := api.NewAPI().
			SetDatabase(dbConn).
			SetPort(":8888").
			SetMaxProcess(1).
			SetSearchClient(searchClient)

		myAPI.GenerateRoute()
	}

}

func getConfig(key string, fallback string) (config string) {
	config = os.Getenv(key)
	if config == "" {
		config = fallback
	}
	return
}
