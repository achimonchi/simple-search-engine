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
	perPage := 100
	resp, err := p.client.
		Collection(INDEX_PRODUCTS).
		Documents().
		Search(&api.SearchCollectionParams{
			Q:       keyword,
			PerPage: &perPage,
			QueryBy: "name,description,category",
		})

	if *resp.Found > 0 {
		hits := resp.Hits
		for _, hit := range *hits {
			var product = ProductEntity{}

			// get document
			h := *hit.Document

			// step 1
			// convert id (string) into integer
			idInt, _ := strconv.Atoi(fmt.Sprintf("%v", h["id"]))
			h["id"] = idInt

			// step 2
			// parse document into product entity
			byteData, err := json.Marshal(h)
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
