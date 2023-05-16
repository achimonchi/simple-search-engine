package api

import (
	"meilisearch/pkg/db"
	"meilisearch/pkg/search"
	"meilisearch/usecase/api/products"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

type API struct {
	router       fiber.Router
	dbConn       db.DatabaseConnection
	searchClient search.Search
	app          *fiber.App
	port         string
	maxProcess   int
}

func NewAPI() API {
	app := fiber.New(fiber.Config{
		AppName: "Search Engine - NBID",
	})

	router := app.Group("/v1")

	// set healthcheck
	router.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "running",
		})
	})

	return API{
		router:     router,
		app:        app,
		maxProcess: runtime.NumCPU(),
	}
}

func (a API) SetDatabase(dbConn db.DatabaseConnection) API {
	a.dbConn = dbConn
	return a
}

func (a API) SetSearchClient(searchClient search.Search) API {
	a.searchClient = searchClient
	return a
}

func (a API) SetPort(port string) API {
	a.port = port
	return a
}
func (a API) SetMaxProcess(maxProcess int) API {
	a.maxProcess = maxProcess
	return a
}

func (a API) GenerateRoute() {
	runtime.GOMAXPROCS(a.maxProcess)

	products.RegisterRoute(a.router, a.dbConn, a.searchClient)
	err := a.app.Listen(a.port)
	if err != nil {
		panic(err)
	}

}
