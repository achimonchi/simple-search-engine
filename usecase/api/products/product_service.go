package products

import (
	"context"
)

// i will seperate the repository based on functionallity

// focus for handle read product
type ProductRead interface {
	GetProductAll(ctx context.Context) (products []ProductEntity, err error)
	GetProductDetailById(ctx context.Context, id int) (product ProductEntity, err error)
}

// focus for handle searching product
type ProductSearch interface {
	SearchProduct(ctx context.Context, keyword string) (products []ProductEntity, err error)
}

// focus for handle write product into database
type ProductWrite interface {
	InsertProduct(ctx context.Context, req ProductModel) (err error)
}

type ProductSearchAndWrite interface {
	ProductSearch
	ProductWrite
}

type ProductReadAndWrite interface {
	ProductRead
	ProductWrite
}

type ProductService struct {
	search ProductSearchAndWrite
	repo   ProductReadAndWrite
}

func NewProductService() ProductService {
	return ProductService{}
}

func (p ProductService) SetRepository(repo ProductReadAndWrite) ProductService {
	p.repo = repo
	return p
}
func (p ProductService) SetSearchRepository(search ProductSearchAndWrite) ProductService {
	p.search = search
	return p
}

func (p ProductService) CreateProduct(ctx context.Context, req ProductModel) (err error) {
	// insert into repository
	err = p.repo.InsertProduct(ctx, req)

	return

}
