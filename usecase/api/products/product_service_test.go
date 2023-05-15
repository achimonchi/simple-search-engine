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

	searchClient := search.NewSearchEngine()

	client, _ := search.ConnectMeili(search.SearchOption{
		Host:   "http://localhost:7700",
		APIKey: "ThisIsMasterKey",
	})
	clientTypesense, _ := search.ConnectTypesense(search.SearchOption{
		Host:   "http://localhost:8108",
		APIKey: "ThisIsMasterKey",
	})

	searchClient = searchClient.SetMeilisearch(client).SetTypesense(clientTypesense)

	builderRepo := NewProductRepository().
		SetDatabaseConnection(dbConn).
		SetSearchEngineClient(searchClient)

	repo := builderRepo.BuildProductRepositoryPostgres()
	searchMeili := builderRepo.BuildProductRepositoryMeilisearch()
	searchTypesense := builderRepo.BuildProductRepositoryTypesense()

	svc = NewProductService().
		SetRepository(repo).
		SetSearchMeiliRepository(searchMeili).
		SetSearchTypesenseRepository(searchTypesense)
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

func TestSearchProductByTypesense(t *testing.T) {
	ctx := context.Background()

	products, err := svc.SearchProductTypesense(ctx, "cooper")
	require.Nil(t, err)
	fmt.Println(products)
}

func BenchmarkSearchProductByMeilisearch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		products, err := svc.SearchProduct(ctx, "cooper")
		_ = products
		_ = err
	}
}
func BenchmarkSearchProductByTypesense(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		products, err := svc.SearchProductTypesense(ctx, "cooper")
		_ = products
		_ = err
	}
}
