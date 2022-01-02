package products

import (
	_productsDomain "acp-final/business/products"
	_controllers "acp-final/controllers"
	_request "acp-final/controllers/products/request"
	_response "acp-final/controllers/products/response"
	"fmt"

	"strconv"

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
	c.Bind(&productRequest) // error handling

	value := c.Param("id")
	number, err := strconv.ParseUint(value, 10, 32)

	finalNumber := _request.ProductId{
		Id: uint(number),
	}

	Product, err := controller.usecase.GetProductById(ctx, int(finalNumber.Id))

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(Product))

}

func (controller *ProductController) GetProductByCategoryId(c echo.Context) error {
	var productRequest _request.CategoryId
	ctx := c.Request().Context()
	if err := c.Bind(&productRequest); err != nil {
		return err
	}

	value := c.Param("category_id")
	number, err := strconv.ParseUint(value, 10, 32)

	finalNumber := _request.CategoryId{
		CategoryId: uint(number),
	}

	products, err := controller.usecase.GetProductByCategoryId(ctx, finalNumber.CategoryId)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(products))
}

func (controller *ProductController) CreateProduct(c echo.Context) error {
	var productRequest _request.CreateProduct
	ctx := c.Request().Context()
	if err := c.Bind(&productRequest); err != nil {
		return err
	}

	dataDomain := _productsDomain.ProductDomain{
		Name:        productRequest.Name,
		CategoryId:  uint(productRequest.CategoryId),
		Description: productRequest.Description,
		Price:       int(productRequest.Price),
		Stock:       int(productRequest.Stock),
	}

	product, err := controller.usecase.CreateProduct(ctx, dataDomain)

	if err != nil {

		return _controllers.NewErrorResponse(c, err)

	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(product))

}

func (controller *ProductController) UpdateProduct(c echo.Context) error {
	var productRequest _request.UpdateProduct
	ctx := c.Request().Context()
	if err := c.Bind(&productRequest); err != nil {
		return err
	}

	dataDomain := _productsDomain.ProductDomain{
		Id:          productRequest.Id,
		Name:        productRequest.Name,
		CategoryId:  uint(productRequest.CategoryId),
		Description: productRequest.Description,
		Price:       int(productRequest.Price),
		Stock:       int(productRequest.Stock),
	}

	fmt.Println(dataDomain)

	product, err := controller.usecase.UpdateProduct(ctx, dataDomain)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(product))

}

func (controller *ProductController) DeleteProduct(c echo.Context) error {
	var productRequest _request.DeleteProduct
	ctx := c.Request().Context()
	if err := c.Bind(&productRequest); err != nil {
		return err
	}

	Product, err := controller.usecase.DeleteProduct(ctx, int(productRequest.Id))

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(Product))
}
