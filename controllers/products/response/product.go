package response

import (
	_productsDomain "acp-final/business/products"
	"time"
)

type ProductResponse struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryId  int       `json:"category_id"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(product _productsDomain.ProductDomain) ProductResponse {
	return ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		CategoryId:  product.CategoryId,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ToListFromDomain(products []_productsDomain.ProductDomain) []ProductResponse {
	var result = []ProductResponse{}
	for _, user := range products {
		result = append(result, FromDomain(user))
	}
	return result
}
