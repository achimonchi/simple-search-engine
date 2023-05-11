package products

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/meilisearch/meilisearch-go"
)

func NewProductRepositoryMeilisearch(client *meilisearch.Client) ProductRepositoryMeilisearch {
	return ProductRepositoryMeilisearch{
		client: client,
	}
}

type ProductRepositoryMeilisearch struct {
	client *meilisearch.Client
}

// SyncProduct implements ProductSearchAndWrite
func (p ProductRepositoryMeilisearch) SyncProduct(ctx context.Context, products []ProductEntity) (err error) {
	var meiliReqs = []map[string]interface{}{}

	jsonReq, err := json.Marshal(products)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonReq, &meiliReqs)
	if err != nil {
		return
	}

	_, err = p.client.Index("products").AddDocuments(meiliReqs)
	if err != nil {
		return
	}
	return
}

// SearchProduct implements ProductSearchAndWrite
func (p ProductRepositoryMeilisearch) SearchProduct(ctx context.Context, keyword string) (products []ProductEntity, err error) {
	resp, err := p.client.Index("products").Search(keyword, &meilisearch.SearchRequest{
		Limit: 10,
	})
	if err != nil {
		return
	}

	fmt.Println(resp.Hits...)
	return
}

// InsertProduct implements ProductSearchAndWrite
func (p ProductRepositoryMeilisearch) InsertProduct(ctx context.Context, req ProductModel) (lastId int, err error) {

	return
}
