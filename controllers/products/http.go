package products

import (
	_productsDomain "acp-final/business/products"
	_controllers "acp-final/controllers"
	_response "acp-final/controllers/products/response"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	usecase _productsDomain.ProductUseCaseDomain
}

func NewProductController(ProductUsecase _productsDomain.ProductUseCaseDomain) *ProductController {
	return &ProductController{
		usecase: ProductUsecase,
	}
}

func (controller *ProductController) GetAllProducts(c echo.Context) error {
	ctx := c.Request().Context()
	products, err := controller.usecase.GetAllProducts(ctx)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(products))
}

func (controller *ProductController) DeleteProduct(c echo.Context, product_id int) error {
	ctx := c.Request().Context()

	Product, err := controller.usecase.DeleteProduct(ctx, product_id)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(Product))
}
