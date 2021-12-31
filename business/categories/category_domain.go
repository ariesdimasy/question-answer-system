package categories

import (
	"context"
	"time"
)

type CategoryDomain struct {
	Id           uint      `json:"id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CategoryUseCaseDomain interface {
	GetAllCategories(ctx context.Context) ([]CategoryDomain, error)
	GetCategoryById(ctx context.Context, id int) (CategoryDomain, error)
	CreateCategory(ctx context.Context, categoryName string) (CategoryDomain, error)
	UpdateCategory(ctx context.Context, categoryUpdate CategoryDomain) (CategoryDomain, error)
	DeleteCategory(ctx context.Context, id int) (CategoryDomain, error)
}

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) (
		res []CategoryDomain, err error)
	GetCategoryById(ctx context.Context, id int) (res CategoryDomain, err error)
	CreateCategory(ctx context.Context, categoryName string) (res CategoryDomain, err error)
	UpdateCategory(ctx context.Context, categoryUpdate CategoryDomain) (res CategoryDomain, err error)
	DeleteCategory(ctx context.Context, id int) (res CategoryDomain, err error)
}
