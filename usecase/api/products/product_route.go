package products

import (
	"meilisearch/pkg/db"
	"meilisearch/pkg/search"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(router fiber.Router, dbConn db.DatabaseConnection, searchClient search.Search) {
	productRepo := NewProductRepository().SetDatabaseConnection(dbConn).SetSearchEngineClient(searchClient)

	repo := productRepo.BuildProductRepositoryPostgres()
	searchMeili := productRepo.BuildProductRepositoryMeilisearch()
	searchTypesense := productRepo.BuildProductRepositoryTypesense()

	pService := NewProductService().
		SetRepository(repo).
		SetSearchMeiliRepository(searchMeili).
		SetSearchTypesenseRepository(searchTypesense)
	pHandler := NewProductHandler(pService)

	productRoute := router.Group("/products")
	{
		productRoute.Get("/", pHandler.GetAll)
		productRoute.Post("/", pHandler.CreateNewProduct)
		productRoute.Post("/search/meili", pHandler.SearchProductMeili)
		productRoute.Post("/search/typesense", pHandler.SearchProductTypesense)
	}
}
