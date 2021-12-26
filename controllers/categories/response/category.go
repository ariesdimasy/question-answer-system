package response

import (
	_categoriesDomain "acp-final/business/categories"
	"time"
)

type CategoryResponse struct {
	Id           uint      `json:"id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(category _categoriesDomain.CategoryDomain) CategoryResponse {
	return CategoryResponse{
		Id:           category.Id,
		CategoryName: category.CategoryName,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
	}
}

func ToListFromDomain(categories []_categoriesDomain.CategoryDomain) []CategoryResponse {
	var result = []CategoryResponse{}
	for _, user := range categories {
		result = append(result, FromDomain(user))
	}
	return result
}
