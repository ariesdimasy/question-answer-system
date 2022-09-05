package routes

import (
	_users "qa-system/controllers/users"

	// "qa-system/helpers"

	// "qa-system/controllers/products/response"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/gommon/email"
)

type ControllerList struct {
	UserController _users.UserController
}

func (cl *ControllerList) RouteRegister(c *echo.Echo) {
	c.GET("/users", cl.UserController.GetAllUser)
	c.POST("/users/login", cl.UserController.Login)
	c.POST("/users/register", cl.UserController.Register)

}
