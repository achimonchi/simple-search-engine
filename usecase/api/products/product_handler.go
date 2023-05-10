package products

import "github.com/gofiber/fiber/v2"

type ProductHandler struct {
	svc ProductService
}

func NewProductHandler(svc ProductService) ProductHandler {
	return ProductHandler{
		svc: svc,
	}
}

func (p ProductHandler) GetAll(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"path": "getall",
	})
}
