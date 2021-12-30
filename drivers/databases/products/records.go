package products

import (
	_productDomain "acp-final/business/products"
	"time"
)

type Product struct {
	Id          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryId  uint      `json:"category_id"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (product *Product) ToDomain() _productDomain.ProductDomain {
	return _productDomain.ProductDomain{
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

func ToListDomain(products []Product) []_productDomain.ProductDomain {
	var result = []_productDomain.ProductDomain{}
	for _, product := range products {
		result = append(result, product.ToDomain())
	}
	return result
}
