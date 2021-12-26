package entities

import (
	"context"
	"time"
)

type ProductDomain struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryId  int       `json:"category_id"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CategoryCreateDomain struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  int    `json:"category_id"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type CategoryUpdateDomain struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  int    `json:"category_id"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type ProductUseCase interface {
	GetAllProducts(ctx context.Context) ([]ProductDomain, error)
	GetProductById(ctx context.Context, id int) (ProductDomain, error)
	GetByCategoryId(ctx context.Context, category_id int) ([]ProductDomain, error)
	CreateProduct(ctx context.Context, createProduct CategoryCreateDomain) (ProductDomain, error)
	UpdateProduct(ctx context.Context, updateProduct CategoryUpdateDomain) (ProductDomain, error)
}

type ProductRepository interface {
	GetAllProducts(ctx context.Context) (
		res []ProductDomain, err error)
	GetById(ctx context.Context, id int) (ProductDomain, error)
	GetByCategoryId(ctx context.Context, category_id int) ([]ProductDomain, error)
	CreateProduct(ctx context.Context, createProduct CategoryCreateDomain) (ProductDomain, error)
	UpdateProduct(ctx context.Context, updateProduct CategoryUpdateDomain) (ProductDomain, error)
}
