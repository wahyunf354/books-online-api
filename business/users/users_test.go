package users_test

import (
	"books_online_api/app/middlewares"
	"books_online_api/business/users"
	_mockUserRepository "books_online_api/business/users/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var userRepository _mockUserRepository.Repository
var configJwt middlewares.ConfigJwt

var userService users.Usecase
var userDomain users.Domain

func setupLogin() {
	userService = users.NewUserUseCase(&userRepository, time.Hour*1, configJwt)
	userDomain = users.Domain{
		Id:           1,
		FirstName:    "Wahyu",
		LastName:     "Fadil",
		Email:        "wahyu@wahyu.com",
		Role:         3,
		Birth:        time.Now(),
		Gender:       "L",
		HashPassword: "$2a$04$Nouh3qsvQ4/CdUj8RZ6SSuX/Vr/gYZcned1Fnucbz9LE888wQ04HW",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}

func setupRegister() {
	userService = users.NewUserUseCase(&userRepository, time.Hour*1, configJwt)
	userDomain = users.Domain{
		Id:           1,
		FirstName:    "Wahyu",
		LastName:     "Fadil",
		Email:        "wahyu@gmail.com",
		Role:         2,
		Birth:        time.Time{},
		Gender:       "L",
		Password:     "123",
		ConfPassword: "123",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}

func TestUserUseCase_Login(t *testing.T) {

	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		setupLogin()
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		user, err := userService.Login(context.Background(), "wahyu@gmail.com", "123")
		assert.Nil(t, err)
		assert.Equal(t, "Wahyu", user.FirstName)
	})

	t.Run("Test Case 2 | Email Empty", func(t *testing.T) {
		setupLogin()
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		_, err := userService.Login(context.Background(), "", "")
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("email empty"),  err)
	})

	t.Run("Test Case 3 | Invalid Email", func(t *testing.T) {
		setupLogin()
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		_, err := userService.Login(context.Background(), "wahyu", "123")
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("email not valid"),  err)
	})


	t.Run("Test Case 4 | Password Empty", func(t *testing.T) {
		setupLogin()
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()


		_, err := userService.Login(context.Background(), "wahyu@gmail.com", "")
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("password empty"),  err)
	})

	t.Run("Test Case 5 | Password Wrong", func(t *testing.T) {
		setupLogin()
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		_, err := userService.Login(context.Background(), "wahyu@gmail.com", "12")
		assert.NotNil(t, err)
		assert.Equal(t, errors.New("password wrong"),  err)
	})

}

func TestUserUseCase_Register(t *testing.T) {
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		setupRegister()

		userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		user, err := userService.Register(context.Background(), users.Domain{
			FirstName:    "Wahyu",
			LastName:     "Fadil",
			Email:        "wahyu@gmail.com",
			Role:         2,
			Birth:        time.Now(),
			Gender:       "L",
			Password:     "123",
			ConfPassword: "123",
		})
		assert.Nil(t, err)
		assert.Equal(t, "Wahyu", user.FirstName)
	})

	t.Run("Test Case 2 | First Name Empty", func(t *testing.T) {
		setupRegister()

		userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		user, err := userService.Register(context.Background(), users.Domain{
			FirstName:    "",
			LastName:     "Fadil",
			Email:        "wahyu@gmail.com",
			Role:         2,
			Birth:        time.Now(),
			Gender:       "L",
			Password:     "123",
			ConfPassword: "123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "", user.FirstName)
		assert.Equal(t, errors.New("first name empty"),  err)
	})


	t.Run("Test Case 3 | Email Empty", func(t *testing.T) {
		setupRegister()

		userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		user, err := userService.Register(context.Background(), users.Domain{
			FirstName:    "Wahyu",
			LastName:     "Fadil",
			Email:        "",
			Role:         2,
			Birth:        time.Now(),
			Gender:       "L",
			Password:     "123",
			ConfPassword: "123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "", user.Email)
		assert.Equal(t, errors.New("email empty"),  err)
	})


	t.Run("Test Case 4 | Invalid Email", func(t *testing.T) {
		setupRegister()

		userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		user, err := userService.Register(context.Background(), users.Domain{
			FirstName:    "Wahyu",
			LastName:     "Fadil",
			Email:        "wahyu",
			Role:         2,
			Birth:        time.Now(),
			Gender:       "L",
			Password:     "123",
			ConfPassword: "123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "", user.Email)
		assert.Equal(t, errors.New("email not valid"),  err)
	})

	t.Run("Test Case 5 | Last name empty", func(t *testing.T) {
		setupRegister()

		userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		user, err := userService.Register(context.Background(), users.Domain{
			FirstName:    "Wahyu",
			LastName:     "",
			Email:        "wahyu@gmail.com",
			Role:         2,
			Birth:        time.Now(),
			Gender:       "L",
			Password:     "123",
			ConfPassword: "123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "", user.LastName)
		assert.Equal(t, errors.New("last name empty"),  err)
	})

	t.Run("Test Case 6 | Password empty", func(t *testing.T) {
		setupRegister()

		userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		user, err := userService.Register(context.Background(), users.Domain{
			FirstName:    "Wahyu",
			LastName:     "Fadil",
			Email:        "wahyu@gmail.com",
			Role:         2,
			Birth:        time.Now(),
			Gender:       "L",
			Password:     "",
			ConfPassword: "123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "", user.Password)
		assert.Equal(t, errors.New("password empty"),  err)
	})

	t.Run("Test Case 7 | Conf Password empty", func(t *testing.T) {
		setupRegister()

		userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		user, err := userService.Register(context.Background(), users.Domain{
			FirstName:    "Wahyu",
			LastName:     "Fadil",
			Email:        "wahyu@gmail.com",
			Role:         2,
			Birth:        time.Now(),
			Gender:       "L",
			Password:     "123",
			ConfPassword: "",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "", user.ConfPassword)
		assert.Equal(t, errors.New("confirm password empty"),  err)
	})

	t.Run("Test Case 8 | Password Not Same", func(t *testing.T) {
		setupRegister()

		userRepository.On("Register", mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		user, err := userService.Register(context.Background(), users.Domain{
			FirstName:    "Wahyu",
			LastName:     "Fadil",
			Email:        "wahyu@gmail.com",
			Role:         2,
			Birth:        time.Now(),
			Gender:       "L",
			Password:     "123",
			ConfPassword: "12",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "", user.ConfPassword)
		assert.Equal(t, "", user.Password)
		assert.Equal(t, errors.New("password not same"),  err)
	})


}