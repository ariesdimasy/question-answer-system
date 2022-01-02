package products

import (
	_productDomain "acp-final/business/products"
	"acp-final/helpers"
	"context"

	"gorm.io/gorm"
)

type ProductRepositoryDatabase struct {
	db *gorm.DB
}

func NewProductRepository(database *gorm.DB) _productDomain.ProductRepository {
	return &ProductRepositoryDatabase{
		db: database,
	}
}

func (repo *ProductRepositoryDatabase) GetAllProducts(ctx context.Context) ([]_productDomain.ProductDomain, error) {

	var products []Product
	result := repo.db.Find(&products)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_productDomain.ProductDomain{}, helpers.ErrNotFound
		} else {
			return []_productDomain.ProductDomain{}, helpers.ErrDbServer
		}
	}

	return ToListDomain(products), nil
}

func (repo *ProductRepositoryDatabase) GetProductByCategoryId(ctx context.Context, category_id uint) ([]_productDomain.ProductDomain, error) {
	var products []Product
	result := repo.db.Where("category_id = ?", category_id).Find(&products)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_productDomain.ProductDomain{}, helpers.ErrNotFound
		} else {
			return []_productDomain.ProductDomain{}, helpers.ErrDbServer
		}
	}

	return ToListDomain(products), nil
}

func (repo *ProductRepositoryDatabase) GetProductById(ctx context.Context, product_id int) (_productDomain.ProductDomain, error) {
	var productResult Product
	result := repo.db.FirstOrInit(&productResult, map[string]interface{}{"id": product_id})

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _productDomain.ProductDomain{}, helpers.ErrNotFound
		} else {
			return _productDomain.ProductDomain{}, helpers.ErrDbServer
		}
	}

	return productResult.ToDomain(), nil
}

func (repo *ProductRepositoryDatabase) CreateProduct(ctx context.Context, createProduct _productDomain.ProductDomain) (_productDomain.ProductDomain, error) {
	var productResult Product = Product{
		Name:        createProduct.Name,
		Description: createProduct.Description,
		CategoryId:  createProduct.CategoryId,
		Stock:       createProduct.Stock,
		Price:       createProduct.Price,
	}

	result := repo.db.Create(&productResult) // return full row { ID:1, Name:"asdas"}

	// repo.db.Where("id = ?", &createProduct.Id).Find(&productResult)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _productDomain.ProductDomain{}, helpers.ErrNotFound
		} else {
			return _productDomain.ProductDomain{}, helpers.ErrDbServer
		}
	}

	return productResult.ToDomain(), nil

}

func (repo *ProductRepositoryDatabase) UpdateProduct(ctx context.Context, updateProduct _productDomain.ProductDomain) (_productDomain.ProductDomain, error) {
	var productResult Product = Product{
		Name:        updateProduct.Name,
		Description: updateProduct.Description,
		CategoryId:  updateProduct.CategoryId,
		Stock:       updateProduct.Stock,
		Price:       updateProduct.Price,
	}

	result := repo.db.Model(&Product{}).Where("id = ?", updateProduct.Id).Updates(productResult).Find(&Product{})

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _productDomain.ProductDomain{}, helpers.ErrNotFound
		} else {
			return _productDomain.ProductDomain{}, helpers.ErrDbServer
		}
	}

	return productResult.ToDomain(), nil

}

func (repo *ProductRepositoryDatabase) DeleteProduct(ctx context.Context, product_id int) (_productDomain.ProductDomain, error) {
	var productResult Product
	repo.db.Where("id = ?", &product_id).Find(&productResult)

	result := repo.db.Delete(&Product{}, product_id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _productDomain.ProductDomain{}, helpers.ErrNotFound
		} else {
			return _productDomain.ProductDomain{}, helpers.ErrDbServer
		}
	}

	return productResult.ToDomain(), nil

}
