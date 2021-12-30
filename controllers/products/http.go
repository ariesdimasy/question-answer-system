package products

import (
	_productsDomain "acp-final/business/products"
	_controllers "acp-final/controllers"
	_request "acp-final/controllers/products/request"
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

func (controller *ProductController) GetProductById(c echo.Context) error {
	ctx := c.Request().Context()
	productRequest := _request.ProductId{}

	Product, err := controller.usecase.GetProductById(ctx, int(productRequest.Id))

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(Product))

}

func (controller *ProductController) GetProductByCategoryId(c echo.Context) error {
	ctx := c.Request().Context()
	productRequest := _request.CategoryId{}

	products, err := controller.usecase.GetProductByCategoryId(ctx, productRequest.CategoryId)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(products))
}

func (controller *ProductController) CreateProduct(c echo.Context) error {
	ctx := c.Request().Context()
	productRequest := _request.CreateProduct{}

	product, err := controller.usecase.CreateProduct(ctx, _productsDomain.ProductCreateDomain{
		Id:          productRequest.Id,
		Name:        productRequest.Name,
		Description: productRequest.Description,
		CategoryId:  productRequest.CategoryId,
		Stock:       productRequest.Stock,
		Price:       productRequest.Price,
	})

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(product))

}

func (controller *ProductController) UpdateProduct(c echo.Context, updateProduct _productsDomain.ProductUpdateDomain) error {
	ctx := c.Request().Context()
	productRequest := _request.UpdateProduct{}

	product, err := controller.usecase.UpdateProduct(ctx, _productsDomain.ProductUpdateDomain{
		Id:          productRequest.Id,
		Name:        productRequest.Name,
		Description: productRequest.Description,
		CategoryId:  productRequest.CategoryId,
		Stock:       productRequest.Stock,
		Price:       productRequest.Price,
	})

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(product))

}

func (controller *ProductController) DeleteProduct(c echo.Context, product_id int) error {
	ctx := c.Request().Context()

	Product, err := controller.usecase.DeleteProduct(ctx, product_id)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(Product))
}
