package products

import (
	"context"
	"time"
)

type ProductDomain struct {
	Id          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryId  int       `json:"category_id"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductCreateDomain struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  int    `json:"category_id"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type ProductUpdateDomain struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  int    `json:"category_id"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type ProductUseCaseDomain interface {
	GetAllProducts(ctx context.Context) ([]ProductDomain, error)
	GetProductById(ctx context.Context, id int) (ProductDomain, error)
	GetProductByCategoryId(ctx context.Context, category_id int) ([]ProductDomain, error)
	CreateProduct(ctx context.Context, createProduct ProductCreateDomain) (ProductDomain, error)
	UpdateProduct(ctx context.Context, updateProduct ProductUpdateDomain) (ProductDomain, error)
	DeleteProduct(ctx context.Context, id int) (ProductDomain, error)
}

type ProductRepository interface {
	GetAllProducts(ctx context.Context) (
		res []ProductDomain, err error)
	GetProductById(ctx context.Context, id int) (res ProductDomain, err error)
	GetProductByCategoryId(ctx context.Context, category_id int) (res []ProductDomain, err error)
	CreateProduct(ctx context.Context, createProduct ProductCreateDomain) (res ProductDomain, err error)
	UpdateProduct(ctx context.Context, updateProduct ProductUpdateDomain) (res ProductDomain, err error)
	DeleteProduct(ctx context.Context, id int) (res ProductDomain, err error)
}
