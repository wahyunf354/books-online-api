package controllers

import (
	"books_online_api/configs"
	"books_online_api/helpers"
	"books_online_api/middlewares"
	"books_online_api/models/response"
	"books_online_api/models/users"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LoginUser(c echo.Context) error {
	userLogin := new(users.UserLogin)

	// Melakukan Bind Data
	if err := c.Bind(userLogin); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Terjadi masalah pada request",
			Data:    err,
		})
	}

	fmt.Println(userLogin)

	user := users.User{}

	result := configs.DB.Where("email = ? && password = ?", userLogin.Email, userLogin.Password).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, response.BaseResponse{
				Code:    http.StatusForbidden,
				Message: "User tidak ditemukan atau password tidak sesuai",
				Data:    nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Ada keselahan di server",
				Data:    result.Error,
			})
		}
	}

	token, err := middlewares.CreateToken(int(user.ID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Ada keselahan di server sini",
			Data:    err,
		})
	}

	userResponse := users.UserResponse{
		Id:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Token:     token,
		Gender:    user.Gender,
		Role:      user.Role,
		Birth:     user.Birth,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil Login",
		Data:    userResponse,
	})
}

func RegisterUser(c echo.Context) error {
	u := new(users.UserRegister)

	// Melakukan Bind Data
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Terjadi masalah pada request",
			Data:    err,
		})
	}

	// validasi jika field kosong
	if u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Role == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Field email, first_name, last_bane, role tidak boleh kosong",
			Data:    nil,
		})
	}

	// validasi password
	if u.Password != u.ConfPassword {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password tidak sama dengan conf password",
			Data:    nil,
		})
	}

	// melakukan encripi password
	hashPassword, err := helpers.HashPassword(&u.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Terjadi masalah saat hash password",
			Data:    err,
		})
	}

	// Melakukan Input Data ke DB
	layout := "2006-01-02T15:04:05.000Z"
	birthUser, _ := time.Parse(layout, u.Birth)

	userNew := users.User{}
	userNew.FirstName = u.FirstName
	userNew.LastName = u.LastName
	userNew.Email = u.Email
	userNew.Gender = u.Gender
	userNew.Role = u.Role
	userNew.Password = hashPassword
	userNew.Birth = birthUser

	result := configs.DB.Create(&userNew)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Terjadi masalah saat input data ke DB",
			Data:    result.Error,
		})
	}

	return c.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Berhasil register user dengan id",
		Data:    userNew,
	})
}
