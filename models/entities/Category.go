package entities

import (
	"context"
	"time"
)

type Category struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	CategoryName string `json:"category_name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryUseCase interface {
	GetAll(ctx context.Context) ([]Category, error)
	GetById(ctx context.Context, id int) (Category, error)
}

type CategoryRepository interface {
	GetAll(ctx context.Context) (
		res []Category, err error)
	GetById(ctx context.Context, id int) (Category, error)
}
