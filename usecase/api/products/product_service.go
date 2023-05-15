package products

import (
	"context"
	"log"
	"time"
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
	SyncProduct(ctx context.Context, products []ProductEntity) (err error)
}

// focus for handle write product into database
type ProductWrite interface {
	InsertProduct(ctx context.Context, req ProductModel) (lastId int, err error)
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
	searchMeili     ProductSearchAndWrite
	searchTypesense ProductSearchAndWrite
	repo            ProductReadAndWrite
}

func NewProductService() ProductService {
	return ProductService{}
}

func (p ProductService) SetRepository(repo ProductReadAndWrite) ProductService {
	p.repo = repo
	return p
}
func (p ProductService) SetSearchMeiliRepository(meili ProductSearchAndWrite) ProductService {
	p.searchMeili = meili
	return p
}
func (p ProductService) SetSearchTypesenseRepository(typesense ProductSearchAndWrite) ProductService {
	p.searchTypesense = typesense
	return p
}

func (p ProductService) CreateProduct(ctx context.Context, req ProductModel) (err error) {

	req.CreatedAt = time.Now()

	// insert into repository
	lastId, err := p.repo.InsertProduct(ctx, req)
	if err != nil {
		return
	}

	req.Id = lastId

	go func() {
		_, err := p.searchMeili.InsertProduct(ctx, req)
		if err != nil {
			log.Println("error when try to insert to meilisearch with error :", err)
		}
	}()
	go func() {
		_, err := p.searchTypesense.InsertProduct(ctx, req)
		if err != nil {
			log.Println("error when try to insert to typesense with error :", err)
		}
	}()

	return

}

func (p ProductService) GetAllProduct(ctx context.Context) (products []ProductEntity, err error) {
	products, err = p.repo.GetProductAll(ctx)
	return
}

func (p ProductService) SearchProduct(ctx context.Context, keyword string) (products []ProductEntity, err error) {
	products, err = p.searchMeili.SearchProduct(ctx, keyword)
	return
}

func (p ProductService) SearchProductTypesense(ctx context.Context, keyword string) (products []ProductEntity, err error) {
	products, err = p.searchTypesense.SearchProduct(ctx, keyword)
	return
}

// if you want to sync data for search engine, use this.
func (p ProductService) syncToSearchEngine(ctx context.Context, req chan bool, done chan bool) {
	process := <-req

	// fmt.Println(process)

	if process {
		products, err := p.repo.GetProductAll(ctx)
		if err != nil {
			log.Println("error when get all products with error :", err.Error())
			return
		}
		err = p.searchMeili.SyncProduct(ctx, products)
		if err != nil {
			log.Println("error when try to sync product with error :", err.Error())
			return
		}
		close(req)
	}

	done <- true
}
