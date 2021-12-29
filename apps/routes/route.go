package routes

import (
	_products "acp-final/controllers/products"
	_users "acp-final/controllers/users"

	// "acp-final/controllers/auth/request"
	// "acp-final/controllers/auth/response"
	// "acp-final/helpers"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/gommon/email"
)

type ControllerList struct {
	UserController    _users.UserController
	ProductController _products.ProductController
}

func (cl *ControllerList) RouteRegister(c *echo.Echo) {
	c.GET("/users", cl.UserController.GetAllUser)
	c.GET("/users/login", cl.UserController.Login(c))

	c.GET("/products", cl.ProductController.GetAllProducts)
}
