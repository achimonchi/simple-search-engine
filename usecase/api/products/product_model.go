package products

import "time"

type ProductEntity struct {
	Id          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	Stock       int       `db:"stock"`
	Category    string    `db:"category"`
	CreatedAt   time.Time `db:"created_at"`
}

type ProductModel struct {
	Name        string    `db:"name"`
	Price       float64   `db:"price"`
	Stock       int       `db:"stock"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	Category    string    `db:"category"`
}
