package users

import (
	"books_online_api/app/middlewares"
	"books_online_api/controllers"
	"books_online_api/helpers"
	"context"
	"time"
)

type UserUseCase struct {
	JWTConfig middlewares.ConfigJwt
	Repo           Repository
	ContextTimeout time.Duration
}

func NewUserUseCase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJwt) Usecase {
	return &UserUseCase{
		JWTConfig:      configJWT,
		Repo:           repo,
		ContextTimeout: timeout,
	}
}

func (uc *UserUseCase) Register(ctx context.Context, user Domain) (Domain, error) {

	if user.Email == "" {
		return Domain{}, controllers.EMPTY_EMAIL
	}

	if !helpers.IsInvalidEmail(user.Email) {
		return Domain{}, controllers.INVALID_EMAIL
	}

	if user.FirstName == "" {
		return Domain{}, controllers.EMPTY_FIRST_NAME
	}

	if user.LastName == "" {
		return Domain{}, controllers.EMPTY_LAST_NAME
	}

	if user.Password == "" {
		return Domain{}, controllers.EMPTY_PASSWORD
	}

	if user.ConfPassword == "" {
		return Domain{}, controllers.EMPTY_CONF_PASSWORD
	}

	if user.Password != user.ConfPassword {
		return Domain{}, controllers.NOT_SAME_PASSWORD
	}

	var errHash error
	user.HashPassword, errHash = helpers.HashPassword(&user.Password)
	if errHash != nil {
		return Domain{}, errHash
	}

	user, err := uc.Repo.Register(ctx, user)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) Login(ctx context.Context, email string, password string) (Domain, error) {
	if email == "" {
		return Domain{}, controllers.EMPTY_EMAIL
	}

	if !helpers.IsInvalidEmail(email) {
		return Domain{}, controllers.INVALID_EMAIL
	}

	if password == "" {
		return Domain{}, controllers.EMPTY_PASSWORD
	}

	user, err := uc.Repo.Login(ctx, email)

	if err != nil {
		return Domain{}, err
	}

	if !helpers.ComparePassword(user.HashPassword, password) {
		return Domain{}, controllers.WRONG_PASSWORD
	}

	user.Token, err = uc.JWTConfig.GenerateTokenJWT(user.Id, user.Role)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
