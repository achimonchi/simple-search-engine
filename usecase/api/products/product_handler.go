package products

import (
	"log"
	"meilisearch/usecase/api/commons/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	svc ProductService
}

func NewProductHandler(svc ProductService) ProductHandler {
	return ProductHandler{
		svc: svc,
	}
}

func (p ProductHandler) GetAll(c *fiber.Ctx) error {
	products, err := p.svc.GetAllProduct(c.Context())
	if err != nil {
		log.Println("error ", err)
		return response.FiberResponse(c, response.Response{
			Status:  http.StatusInternalServerError,
			Error:   err,
			Message: err.Error(),
		})
	}

	return response.FiberResponse(c, response.Response{
		Status:  http.StatusOK,
		Message: "get product success",
		Data:    products,
	})
}

func (p ProductHandler) CreateNewProduct(c *fiber.Ctx) error {
	var req = ProductModel{}

	if err := c.BodyParser(&req); err != nil {
		log.Println("error ", err)
		return response.FiberResponse(c, response.Response{
			Status:  http.StatusBadRequest,
			Error:   err,
			Message: "invalid request",
		})
	}

	if err := p.svc.CreateProduct(c.Context(), req); err != nil {
		log.Println("error ", err)
		return response.FiberResponse(c, response.Response{
			Status:  http.StatusInternalServerError,
			Error:   err,
			Message: err.Error(),
		})
	}

	return response.FiberResponse(c, response.Response{
		Status:  http.StatusCreated,
		Message: "create product success",
	})
}

func (p ProductHandler) SearchProductMeili(c *fiber.Ctx) error {
	var req = ProductSearchModel{}

	if err := c.BodyParser(&req); err != nil {
		log.Println("error search product meili", err)
		return response.FiberResponse(c, response.Response{
			Status:  http.StatusInternalServerError,
			Error:   err,
			Message: err.Error(),
		})
	}

	products, err := p.svc.SearchProduct(c.Context(), req.Keyword)
	if err != nil {
		log.Println("error search product meili", err)
		return response.FiberResponse(c, response.Response{
			Status:  http.StatusInternalServerError,
			Error:   err,
			Message: err.Error(),
		})
	}

	return response.FiberResponse(c, response.Response{
		Status:  http.StatusOK,
		Message: "search success",
		Data:    products,
	})
}

func (p ProductHandler) SearchProductTypesense(c *fiber.Ctx) error {
	var req = ProductSearchModel{}
	if err := c.BodyParser(&req); err != nil {
		log.Println("error search product typesense ", err)
		return response.FiberResponse(c, response.Response{
			Status:  http.StatusInternalServerError,
			Error:   err,
			Message: err.Error(),
		})
	}

	products, err := p.svc.SearchProductTypesense(c.Context(), req.Keyword)
	if err != nil {
		log.Println("error search product typesense ", err)
		return response.FiberResponse(c, response.Response{
			Status:  http.StatusInternalServerError,
			Error:   err,
			Message: err.Error(),
		})
	}

	return response.FiberResponse(c, response.Response{
		Status:  http.StatusOK,
		Message: "search success",
		Data:    products,
	})
}
