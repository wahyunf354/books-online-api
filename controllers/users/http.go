package users

import (
	"books_online_api/business/users"
	"books_online_api/controllers"
	"books_online_api/controllers/users/requests"
	"books_online_api/controllers/users/responses/login"
	"books_online_api/controllers/users/responses/register"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(userUseCase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (userController UserController) Register(c echo.Context) error {
	userRegister := requests.UserRegister{}
	err := c.Bind(&userRegister)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	user, err := userController.UserUseCase.Register(ctx, userRegister.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusCreated, register.FromDomain(user))
}

func (userController UserController) Login(c echo.Context) error {
	userLogin := requests.UserLogin{}
	err := c.Bind(&userLogin)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	user, err := userController.UserUseCase.Login(ctx, userLogin.ToDomain().Email, userLogin.ToDomain().Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, http.StatusOK, login.FromDomain(user))
}
