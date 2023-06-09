package products

import (
	"context"
	"database/sql"
)

func NewProductRepositoryPostgres(db *sql.DB) ProductRepositoryPostgres {
	return ProductRepositoryPostgres{
		db: db,
	}
}

type ProductRepositoryPostgres struct {
	db *sql.DB
}

// GetProductAll implements ProductReadAndWrite
func (p ProductRepositoryPostgres) GetProductAll(ctx context.Context) (products []ProductEntity, err error) {
	query := `
		SELECT id, category, name, description, price, stock, created_at
		FROM products
	`

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return
	}

	for rows.Next() {
		product := ProductEntity{}
		err = rows.Scan(
			&product.Id,
			&product.Category,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
		)

		if err != nil {
			return
		}

		products = append(products, product)
	}

	return

}

// GetProductDetailById implements ProductReadAndWrite
func (p ProductRepositoryPostgres) GetProductDetailById(ctx context.Context, id int) (product ProductEntity, err error) {
	query := `
		SELECT id, name, category, description, price, stock, created_at
		FROM products
		WHERE id=$1
	`

	row := p.db.QueryRowContext(ctx, query, id)
	err = row.Scan(
		&product.Id,
		&product.Name,
		&product.Category,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
	)

	if err != nil {
		return
	}
	return
}

// InsertProduct implements ProductReadAndWrite
func (p ProductRepositoryPostgres) InsertProduct(ctx context.Context, req ProductModel) (lastId int, err error) {
	query := `
		INSERT INTO products (
			name, description, price, stock, created_at, category
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
		RETURNING id;
	`

	err = p.db.QueryRowContext(ctx, query,
		req.Name,
		req.Description,
		req.Price,
		req.Stock,
		req.CreatedAt,
		req.Category,
	).Scan(&lastId)

	return
}
