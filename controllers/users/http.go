package users

import (
	_usersDomain "acp-final/business/users"
	_controllers "acp-final/controllers"
	_response "acp-final/controllers/users/response"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase _usersDomain.UserUseCaseDomain
}

func NewUserController(userUsecase _usersDomain.UserUseCaseDomain) *UserController {
	return &UserController{
		usecase: userUsecase,
	}
}

func (controller *UserController) GetAllUser(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := controller.usecase.GetAllUser(ctx)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(users))
}

func (controller *UserController) Login(c echo.Context, userlogin _usersDomain.UserLoginDomain) error {
	ctx := c.Request().Context()
	user, err := controller.usecase.Login(ctx, userlogin)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(user))
}

func (controller *UserController) Register(c echo.Context, userRegister _usersDomain.UserRegisterDomain) error {
	ctx := c.Request().Context()
	user, err := controller.usecase.Register(ctx, userRegister)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(user))
}

func (controller *UserController) DeleteUser(c echo.Context, user_id string) error {
	ctx := c.Request().Context()

	user, err := controller.usecase.DeleteUser(ctx, user_id)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(user))
}
