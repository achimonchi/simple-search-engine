package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

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

func (s Search) MigrateUp(filename string) (err error) {
	log.Println("try to migrate search data in file", filename)
	defer func() {
		log.Println("migrate search data success")
	}()
	jsonFile, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listProducts []map[string]interface{}
	err = json.Unmarshal(byteValue, &listProducts)
	if err != nil {
		return err
	}

	resp, err := s.Meilisearch.Index("products_search").AddDocuments(listProducts)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return
}
