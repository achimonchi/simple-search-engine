package products

import (
	"meilisearch/pkg/db"
	"meilisearch/pkg/search"
)

type ProductRepository struct {
	dbConn       db.DatabaseConnection
	searchClient search.Search
}

func NewProductRepository() ProductRepository {
	return ProductRepository{}
}

func (p ProductRepository) SetDatabaseConnection(dbConn db.DatabaseConnection) ProductRepository {
	p.dbConn = dbConn
	return p
}
func (p ProductRepository) SetSearchEngineClient(searchClient search.Search) ProductRepository {
	p.searchClient = searchClient
	return p
}

func (p ProductRepository) BuildProductRepositoryPostgres() ProductRepositoryPostgres {
	return NewProductRepositoryPostgres(p.dbConn.Postgres)
}

func (p ProductRepository) BuildProductRepositoryMeilisearch() ProductRepositoryMeilisearch {
	return NewProductRepositoryMeilisearch(p.searchClient.Meilisearch)
}

func (p ProductRepository) BuildProductRepositoryTypesense() ProductRepositoryTypesense {
	return NewProductRepositoryTypesense(p.searchClient.Typesense)
}
