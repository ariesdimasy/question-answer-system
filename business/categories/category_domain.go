package entities

import (
	"context"
	"time"
)

type CategoryDomain struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CategoryUpdateDomain struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	CategoryName string `json:"category_name"`
}

type CategoryUseCase interface {
	GetAllCategories(ctx context.Context) ([]CategoryDomain, error)
	GetCategoryById(ctx context.Context, id int) (CategoryDomain, error)
	CreateCategory(ctx context.Context, categoryName string)
	UpdateCategory(ctx context.Context, categoryUpdate CategoryUpdateDomain)
	DeleteCategory(ctx context.Context, id int)
}

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) (
		res []CategoryDomain, err error)
	GetCategoryById(ctx context.Context, id int) (CategoryDomain, error)
	CreateCategory(ctx context.Context, categoryName string)
	UpdateCategory(ctx context.Context, categoryUpdate CategoryUpdateDomain)
	DeleteCategory(ctx context.Context, id int)
}
