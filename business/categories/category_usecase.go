package categories

import (
	"context"
	"time"
)

type CategoryUsecase struct {
	contextTimeout time.Duration
	repository     CategoryRepository
}

func NewCategoryUsecase(repo CategoryRepository, timeout time.Duration) CategoryUseCaseDomain {
	return &CategoryUsecase{
		contextTimeout: timeout,
		repository:     repo,
	}
}

func (uc *CategoryUsecase) GetAllCategories(ctx context.Context) ([]CategoryDomain, error) {
	categories, err := uc.repository.GetAllCategories(ctx)
	if err != nil {
		return []CategoryDomain{}, err
	}
	return categories, nil
}

func (uc *CategoryUsecase) GetCategoryById(ctx context.Context, id uint) (CategoryDomain, error) {
	category, err := uc.repository.GetCategoryById(ctx, id)
	if err != nil {
		return CategoryDomain{}, err
	}
	return category, nil
}

func (uc *CategoryUsecase) CreateCategory(ctx context.Context, categoryName string) (CategoryDomain, error) {
	category, err := uc.repository.CreateCategory(ctx, categoryName)
	if err != nil {
		return CategoryDomain{}, err
	}
	return category, nil
}

func (uc *CategoryUsecase) UpdateCategory(ctx context.Context, categoryUpdate CategoryDomain) (CategoryDomain, error) {
	category, err := uc.repository.UpdateCategory(ctx, categoryUpdate)
	if err != nil {
		return CategoryDomain{}, err
	}
	return category, nil
}

func (uc *CategoryUsecase) DeleteCategory(ctx context.Context, id uint) (CategoryDomain, error) {
	category, err := uc.repository.DeleteCategory(ctx, id)
	if err != nil {
		return CategoryDomain{}, err
	}
	return category, nil
}
