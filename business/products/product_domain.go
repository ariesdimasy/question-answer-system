package products

import (
	"context"
	"time"
)

type ProductDomain struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryId  uint      `json:"category_id"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductUseCaseDomain interface {
	GetAllProducts(ctx context.Context) ([]ProductDomain, error)
	GetProductById(ctx context.Context, id int) (ProductDomain, error)
	GetProductByCategoryId(ctx context.Context, category_id uint) ([]ProductDomain, error)
	CreateProduct(ctx context.Context, createProduct ProductDomain) (ProductDomain, error)
	UpdateProduct(ctx context.Context, updateProduct ProductDomain) (ProductDomain, error)
	DeleteProduct(ctx context.Context, id int) (ProductDomain, error)
}

type ProductRepository interface {
	GetAllProducts(ctx context.Context) (
		res []ProductDomain, err error)
	GetProductById(ctx context.Context, id int) (res ProductDomain, err error)
	GetProductByCategoryId(ctx context.Context, category_id uint) (res []ProductDomain, err error)
	CreateProduct(ctx context.Context, createProduct ProductDomain) (res ProductDomain, err error)
	UpdateProduct(ctx context.Context, updateProduct ProductDomain) (res ProductDomain, err error)
	DeleteProduct(ctx context.Context, id int) (res ProductDomain, err error)
}
