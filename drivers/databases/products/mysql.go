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

func NewUserRepository(database *gorm.DB) _productDomain.ProductRepository {
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

func (repo *ProductRepositoryDatabase) CreateProduct(ctx context.Context, createProduct _productDomain.ProductCreateDomain) (_productDomain.ProductDomain, error) {
	var productResult Product

	result := repo.db.Create(&createProduct)

	repo.db.Where("id = ?", &createProduct.Id).Find(&productResult)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _productDomain.ProductDomain{}, helpers.ErrNotFound
		} else {
			return _productDomain.ProductDomain{}, helpers.ErrDbServer
		}
	}

	return productResult.ToDomain(), nil

}

func (repo *ProductRepositoryDatabase) UpdateProduct(ctx context.Context, updateProduct _productDomain.ProductUpdateDomain) (_productDomain.ProductDomain, error) {
	var productResult Product

	result := repo.db.Model(&Product{}).Where("id = ?", updateProduct.Id).Updates(Product{
		Name:        updateProduct.Name,
		CategoryId:  updateProduct.CategoryId,
		Description: updateProduct.Description,
		Price:       updateProduct.Price,
		Stock:       updateProduct.Stock,
	})

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
