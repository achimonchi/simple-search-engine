package products

import (
	"meilisearch/pkg/db"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(router fiber.Router, dbConn db.DatabaseConnection) {
	productRepo := NewProductRepository().SetDatabaseConnection(dbConn)

	repo := productRepo.BuildProductRepositoryPostgres()
	pService := NewProductService().SetRepository(repo)
	pHandler := NewProductHandler(pService)

	productRoute := router.Group("/products")
	{
		productRoute.Get("/", pHandler.GetAll)
	}
}
