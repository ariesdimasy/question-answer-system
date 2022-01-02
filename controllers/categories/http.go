package categorys

import (
	_categoryDomain "acp-final/business/categories"
	_controllers "acp-final/controllers"
	_request "acp-final/controllers/categories/request"
	_response "acp-final/controllers/categories/response"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	usecase _categoryDomain.CategoryUseCaseDomain
}

func NewCategoryController(CategoryUsecase _categoryDomain.CategoryUseCaseDomain) *CategoryController {
	return &CategoryController{
		usecase: CategoryUsecase,
	}
}

func (controller *CategoryController) GetAllCategories(c echo.Context) error {
	ctx := c.Request().Context()
	categories, err := controller.usecase.GetAllCategories(ctx)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(categories))
}

func (controller *CategoryController) GetCategoryById(c echo.Context) error {
	ctx := c.Request().Context()
	categoryRequest := _request.CategoryId{}

	Category, err := controller.usecase.GetCategoryById(ctx, categoryRequest.Id)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(Category))

}

func (controller *CategoryController) CreateCategory(c echo.Context) error {
	ctx := c.Request().Context()
	categoryRequest := _request.CreateCategory{}

	category, err := controller.usecase.CreateCategory(ctx, categoryRequest.CategoryName)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(category))

}

func (controller *CategoryController) UpdateCategory(c echo.Context) error {
	ctx := c.Request().Context()
	categoryRequest := _request.UpdateCategory{}

	category, err := controller.usecase.UpdateCategory(ctx, _categoryDomain.CategoryDomain{
		Id:           categoryRequest.Id,
		CategoryName: categoryRequest.CategoryName,
	})

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(category))

}

func (controller *CategoryController) DeleteCategory(c echo.Context) error {
	ctx := c.Request().Context()
	categoryRequest := _request.DeleteCategory{}

	Category, err := controller.usecase.DeleteCategory(ctx, categoryRequest.Id)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(Category))
}
