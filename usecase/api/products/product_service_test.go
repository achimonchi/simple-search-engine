package products

import (
	"context"
	"database/sql"
	"fmt"
	"meilisearch/pkg/db"
	"meilisearch/pkg/search"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc ProductService

func init() {
	dbConn, _ := db.ConnectPostgres(db.DatabaseOption{
		Host:    "localhost",
		Port:    "6632",
		User:    "user-search",
		Pass:    "user-pass",
		Dbname:  "search",
		Sslmode: db.SSL_DISABLE,
	})

	searchClient, _ := search.ConnectMeili(search.SearchOption{
		Host:   "http://localhost:7700",
		APIKey: "ThisIsMasterKey",
	})

	searchClient.Meilisearch.DeleteIndex("products")

	builderRepo := NewProductRepository().
		SetDatabaseConnection(dbConn).
		SetSearchEngineClient(searchClient)

	repo := builderRepo.BuildProductRepositoryPostgres()
	searchEngine := builderRepo.BuildProductRepositoryMeilisearch()
	svc = NewProductService().
		SetRepository(repo).
		SetSearchRepository(searchEngine)
}

func TestCreateProduct(t *testing.T) {
	ctx := context.Background()
	req := ProductModel{
		Name:        "Product 2",
		Price:       10000,
		Stock:       20,
		Description: "Ini adalah product 1",
		Category:    "Jaket",
	}

	err := svc.CreateProduct(ctx, req)
	require.Nil(t, err)
}

func TestGetAllProduct(t *testing.T) {
	ctx := context.Background()

	products, err := svc.GetAllProduct(ctx)
	if err != nil {
		require.NotEqual(t, err.Error(), sql.ErrNoRows)
	}

	for i, p := range products {
		str := strings.Repeat("=", 4)
		fmt.Println(str, "[", i+1, "]", str)
		fmt.Println("Id \t\t:", p.Id)
		fmt.Println("Name \t\t:", p.Name)
		fmt.Println("Category \t:", p.Category)
		fmt.Println("Description \t:", p.Description)
		fmt.Println("Price \t\t:", p.Price)
		fmt.Println("Stock \t\t:", p.Stock)
		fmt.Println("CreatedAt \t:", p.CreatedAt)
		fmt.Println(str + str)
	}
}

func TestSearchProduct(t *testing.T) {
	ctx := context.Background()

	products, err := svc.SearchProduct(ctx, "jaxket")
	require.Nil(t, err)
	require.NotNil(t, products)
}
