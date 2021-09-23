package users

import (
	"books_online_api/business/users"
	"books_online_api/controllers"
	"books_online_api/controllers/users/requests"
	"books_online_api/controllers/users/responses"
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

func (UserController UserController) Register(c echo.Context) error {
	userRegister := requests.UserRegister{}
	c.Bind(&userRegister)

	ctx := c.Request().Context()
	user, error := UserController.UserUseCase.Register(ctx, userRegister.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, http.StatusCreated, responses.FromDomain(user))
}
