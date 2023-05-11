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
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Category    string    `json:"category"`
}
