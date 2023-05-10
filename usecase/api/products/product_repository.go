package products

import "meilisearch/pkg/db"

type ProductRepository struct {
	dbConn db.DatabaseConnection
}

func NewProductRepository() ProductRepository {
	return ProductRepository{}
}

func (p ProductRepository) SetDatabaseConnection(dbConn db.DatabaseConnection) ProductRepository {
	p.dbConn = dbConn
	return p
}

func (p ProductRepository) BuildProductRepositoryPostgres() ProductRepositoryPostgres {
	return NewProductRepositoryPostgres(p.dbConn.Postgres)
}
