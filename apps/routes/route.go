package routes

import (
	_users "ACP14/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController _users.UserController
}

func (cl *ControllerList) RouteRegister(c *echo.Echo) {
	c.GET("/users", cl.UserController.GetUser)
}
