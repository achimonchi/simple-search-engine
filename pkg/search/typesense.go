package search

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

func ConnectTypesense(option SearchOption) (client *typesense.Client, err error) {
	client = typesense.NewClient(
		typesense.WithServer(option.Host),
		typesense.WithAPIKey(option.APIKey),
		typesense.WithConnectionTimeout(5*time.Second),
	)
	health, err := client.Health(5 * time.Second)
	if !health || err != nil {
		return
	}

	// s.Typesense = client
	return
}

func (s Search) MigrateTypesenseUp(filename string) (err error) {
	log.Println("try to migrate search data for typesense in file", filename)
	defer func() {
		log.Println("migrate search data success")
	}()

	// defaultSortingField := "id"
	schema := &api.CollectionSchema{
		Name: "products_search",
		Fields: []api.Field{
			{
				Name: "id",
				Type: "int32",
			},
			{
				Name: "name",
				Type: "string",
			},
			{
				Name: "description",
				Type: "string",
			},
			{
				Name: "category",
				Type: "string",
			},
			{
				Name: "price",
				Type: "int32",
			},
			{
				Name: "stock",
				Type: "int32",
			},
		},
		// DefaultSortingField: &defaultSortingField,
	}

	_, err = s.Typesense.Collections().Create(schema)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

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

	for _, product := range listProducts {
		// id should be string
		product["id"] = fmt.Sprintf("%v", product["id"])
		resp, err := s.Typesense.Collection("products_search").Documents().Create(product)
		if err != nil {
			return err
		}
		fmt.Printf("%+v\n", resp)
	}

	return
}
