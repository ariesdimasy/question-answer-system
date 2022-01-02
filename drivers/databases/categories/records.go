package categories

import (
	_categoryDomain "acp-final/business/categories"
	"time"
)

type Category struct {
	Id           uint      `json:"id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

func (category *Category) ToDomain() _categoryDomain.CategoryDomain {
	return _categoryDomain.CategoryDomain{
		Id:           category.Id,
		CategoryName: category.CategoryName,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}
}

func ToListDomain(categorys []Category) []_categoryDomain.CategoryDomain {
	var result = []_categoryDomain.CategoryDomain{}
	for _, category := range categorys {
		result = append(result, category.ToDomain())
	}
	return result
}
