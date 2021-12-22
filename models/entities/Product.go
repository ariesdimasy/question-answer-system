package entities

import (
	"context"
	"time"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryId  int       `json:"category_id"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductUseCase interface {
	GetAll(ctx context.Context) ([]Product, error)
	GetById(ctx context.Context, id int) (Product, error)
	GetByCategoryId(ctx context.Context, category_id int) ([]Product, error)
}

type ProductRepository interface {
	GetAll(ctx context.Context) (
		res []Product, err error)
	GetById(ctx context.Context, id int) (Product, error)
	GetByCategoryId(ctx context.Context, category_id int) ([]Product, error)
}
