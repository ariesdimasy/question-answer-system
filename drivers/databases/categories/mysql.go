package categories

import (
	_categoryDomain "acp-final/business/categories"
	"acp-final/helpers"
	"context"

	"gorm.io/gorm"
)

type categoryRepositoryDatabase struct {
	db *gorm.DB
}

func NewCategoryRepository(database *gorm.DB) _categoryDomain.CategoryRepository {
	return &categoryRepositoryDatabase{
		db: database,
	}
}

func (repo *categoryRepositoryDatabase) GetAllCategories(ctx context.Context) ([]_categoryDomain.CategoryDomain, error) {

	var categories []Category
	result := repo.db.Find(&categories)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_categoryDomain.CategoryDomain{}, helpers.ErrNotFound
		} else {
			return []_categoryDomain.CategoryDomain{}, helpers.ErrDbServer
		}
	}

	return ToListDomain(categories), nil
}

func (repo *categoryRepositoryDatabase) GetCategoryById(ctx context.Context, category_id uint) (_categoryDomain.CategoryDomain, error) {
	var category Category
	result := repo.db.Debug().Where("id = ?", category_id).Find(&category)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _categoryDomain.CategoryDomain{}, helpers.ErrNotFound
		} else {
			return _categoryDomain.CategoryDomain{}, helpers.ErrDbServer
		}
	}

	return category.ToDomain(), nil
}

func (repo *categoryRepositoryDatabase) CreateCategory(ctx context.Context, category_name string) (_categoryDomain.CategoryDomain, error) {
	var categoryResult Category = Category{
		CategoryName: category_name,
	}

	result := repo.db.Create(&categoryResult)

	repo.db.Where("id = ?", &categoryResult.Id).Find(&categoryResult)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _categoryDomain.CategoryDomain{}, helpers.ErrNotFound
		} else {
			return _categoryDomain.CategoryDomain{}, helpers.ErrDbServer
		}
	}

	return categoryResult.ToDomain(), nil

}

func (repo *categoryRepositoryDatabase) UpdateCategory(ctx context.Context, updatecategory _categoryDomain.CategoryDomain) (_categoryDomain.CategoryDomain, error) {
	var categoryResult Category = Category{
		Id:           updatecategory.Id,
		CategoryName: updatecategory.CategoryName,
	}

	result := repo.db.Model(&Category{}).Where("id = ?", updatecategory.Id).Updates(Category{
		CategoryName: updatecategory.CategoryName,
	}).Find(&Category{})

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _categoryDomain.CategoryDomain{}, helpers.ErrNotFound
		} else {
			return _categoryDomain.CategoryDomain{}, helpers.ErrDbServer
		}
	}

	return categoryResult.ToDomain(), nil

}

func (repo *categoryRepositoryDatabase) DeleteCategory(ctx context.Context, category_id uint) (_categoryDomain.CategoryDomain, error) {
	var categoryResult Category
	repo.db.Where("id = ?", &category_id).Find(&categoryResult)

	result := repo.db.Delete(&Category{}, category_id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _categoryDomain.CategoryDomain{}, helpers.ErrNotFound
		} else {
			return _categoryDomain.CategoryDomain{}, helpers.ErrDbServer
		}
	}

	return categoryResult.ToDomain(), nil

}
