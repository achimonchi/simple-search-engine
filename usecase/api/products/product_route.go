package products

import (
	"meilisearch/pkg/db"
	"meilisearch/pkg/search"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(router fiber.Router, dbConn db.DatabaseConnection, searchClient search.Search) {
	productRepo := NewProductRepository().SetDatabaseConnection(dbConn).SetSearchEngineClient(searchClient)

	repo := productRepo.BuildProductRepositoryPostgres()
	searchEngine := productRepo.BuildProductRepositoryMeilisearch()
	pService := NewProductService().
		SetRepository(repo).
		SetSearchRepository(searchEngine)
	pHandler := NewProductHandler(pService)

	productRoute := router.Group("/products")
	{
		productRoute.Get("/", pHandler.GetAll)
		productRoute.Post("/", pHandler.CreateNewProduct)
	}
}
