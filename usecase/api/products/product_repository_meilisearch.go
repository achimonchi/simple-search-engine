package products

import (
	"context"
	"encoding/json"

	"github.com/meilisearch/meilisearch-go"
)

func NewProductRepositoryMeilisearch(client *meilisearch.Client) ProductRepositoryMeilisearch {
	return ProductRepositoryMeilisearch{
		client: client,
	}
}

const INDEX_PRODUCTS = "products_search"

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

	_, err = p.client.Index(INDEX_PRODUCTS).UpdateDocuments(meiliReqs)
	if err != nil {
		return
	}
	return
}

// SearchProduct implements ProductSearchAndWrite
func (p ProductRepositoryMeilisearch) SearchProduct(ctx context.Context, keyword string) (products []ProductEntity, err error) {
	resp, err := p.client.Index(INDEX_PRODUCTS).Search(keyword, &meilisearch.SearchRequest{
		Limit: 10,
	})

	if err != nil {
		return
	}

	respByte, err := json.Marshal(resp.Hits)
	if err != nil {
		return
	}

	err = json.Unmarshal(respByte, &products)

	if err != nil {
		return
	}

	return
}

// InsertProduct implements ProductSearchAndWrite
func (p ProductRepositoryMeilisearch) InsertProduct(ctx context.Context, req ProductModel) (lastId int, err error) {
	var meiliReqs = []map[string]interface{}{}

	jsonReq, err := json.Marshal(req)
	if err != nil {
		return
	}

	var newReq = map[string]interface{}{}

	err = json.Unmarshal(jsonReq, &newReq)
	if err != nil {
		return
	}

	newReq["id"] = req.Id

	meiliReqs = append(meiliReqs, newReq)

	_, err = p.client.Index(INDEX_PRODUCTS).UpdateDocuments(meiliReqs)
	if err != nil {
		return
	}
	return
}
