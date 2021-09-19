package controllers

import (
	"books_online_api/configs"
	"books_online_api/models/response"
	"books_online_api/models/users"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

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
	hashPassword, err := HashPassword(&u.Password)
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

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil register user dengan id",
		Data:    userNew,
	})
}

func HashPassword(password *string) (string, error) {
	// convert password string to byte
	passwordBytes := []byte(*password)

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPassword), err
}
