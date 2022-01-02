package products

import (
	"context"
	"time"
)

type ProductUsecase struct {
	contextTimeout time.Duration
	repository     ProductRepository
}

func NewProductUsecase(repo ProductRepository, timeout time.Duration) ProductUseCaseDomain {
	return &ProductUsecase{
		contextTimeout: timeout,
		repository:     repo,
	}
}

func (uc *ProductUsecase) GetAllProducts(ctx context.Context) ([]ProductDomain, error) {
	// timeout
	Products, err := uc.repository.GetAllProducts(ctx)
	if err != nil {
		return []ProductDomain{}, err
	}
	return Products, nil
}

func (uc *ProductUsecase) GetProductById(ctx context.Context, id int) (ProductDomain, error) {
	// timeout
	Product, err := uc.repository.GetProductById(ctx, id)
	if err != nil {
		return ProductDomain{}, err
	}
	return Product, nil
}

func (uc *ProductUsecase) GetProductByCategoryId(ctx context.Context, category_id uint) ([]ProductDomain, error) {
	// timeout
	Product, err := uc.repository.GetProductByCategoryId(ctx, category_id)
	if err != nil {
		return []ProductDomain{}, err
	}
	return Product, nil
}

func (uc *ProductUsecase) CreateProduct(ctx context.Context, createProduct ProductDomain) (ProductDomain, error) {
	Product, err := uc.repository.CreateProduct(ctx, createProduct)
	if err != nil {
		return ProductDomain{}, err
	}
	return Product, nil
}

func (uc *ProductUsecase) UpdateProduct(ctx context.Context, updateProduct ProductDomain) (ProductDomain, error) {
	Product, err := uc.repository.UpdateProduct(ctx, updateProduct)
	if err != nil {
		return ProductDomain{}, err
	}
	return Product, nil
}

func (uc *ProductUsecase) DeleteProduct(ctx context.Context, id int) (ProductDomain, error) {
	Product, err := uc.repository.DeleteProduct(ctx, id)
	if err != nil {
		return ProductDomain{}, err
	}
	return Product, nil
}
