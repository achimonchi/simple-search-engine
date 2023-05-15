package products

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

func NewProductRepositoryTypesense(client *typesense.Client) ProductRepositoryTypesense {
	return ProductRepositoryTypesense{
		client: client,
	}
}

type ProductRepositoryTypesense struct {
	client *typesense.Client
}

// SearchProduct implements ProductSearchAndWrite
func (p ProductRepositoryTypesense) SearchProduct(ctx context.Context, keyword string) (products []ProductEntity, err error) {
	resp, err := p.client.
		Collection(INDEX_PRODUCTS).
		Documents().
		Search(&api.SearchCollectionParams{
			Q:       keyword,
			QueryBy: "name,description",
		})

	if *resp.Found > 0 {
		hits := resp.Hits
		for _, hit := range *hits {
			var product = ProductEntity{}
			var data = map[string]interface{}{}

			// step 1
			// parse into byte from hit data
			byteData, err := json.Marshal(hit)
			if err != nil {
				return products, err
			}

			err = json.Unmarshal(byteData, &data)
			if err != nil {
				return products, err
			}

			// step 2
			// parse document data into map[string]
			byteData, err = json.Marshal(data["document"])
			if err != nil {
				return products, err
			}

			err = json.Unmarshal(byteData, &data)
			if err != nil {
				return products, err
			}

			// step 3
			// convert id (string) into integer
			idInt, _ := strconv.Atoi(fmt.Sprintf("%v", data["id"]))
			data["id"] = idInt

			// step 4
			// parse again data into product entity
			byteData, err = json.Marshal(data)
			if err != nil {
				return products, err
			}

			err = json.Unmarshal(byteData, &product)
			if err != nil {
				return products, err
			}
			products = append(products, product)
		}
	}
	return
}

// SyncProduct implements ProductSearchAndWrite
func (p ProductRepositoryTypesense) SyncProduct(ctx context.Context, products []ProductEntity) (err error) {
	panic("unimplemented")
}

// InsertProduct implements ProductSearchAndWrite
func (p ProductRepositoryTypesense) InsertProduct(ctx context.Context, req ProductModel) (lastId int, err error) {
	var product map[string]interface{}

	byteProduct, err := json.Marshal(req)
	if err != nil {
		return
	}

	err = json.Unmarshal(byteProduct, &product)
	if err != nil {
		return
	}

	// id must be string
	product["id"] = uuid.NewString()

	_, err = p.client.Collection(INDEX_PRODUCTS).Documents().Upsert(product)
	return
}
