package users

import (
	"books_online_api/app/middlewares"
	"books_online_api/helpers"
	"context"
	"errors"
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
		return Domain{}, errors.New("email empty")
	}

	if !helpers.IsInvalidEmail(user.Email) {
		return Domain{}, errors.New("email not valid")
	}

	if user.FirstName == "" {
		return Domain{}, errors.New("first name empty")
	}

	if user.LastName == "" {
		return Domain{}, errors.New("last name empty")
	}

	if user.Password == "" {
		return Domain{}, errors.New("password empty")
	}

	if user.ConfPassword == "" {
		return Domain{}, errors.New("confirm password empty")
	}

	if user.Password != user.ConfPassword {
		return Domain{}, errors.New("password not same")
	}

	var errHash error
	user.HashPassword, errHash = helpers.HashPassword(&user.Password)
	if errHash != nil {
		return Domain{}, errors.New("faild hashing password")
	}

	user, err := uc.Repo.Register(ctx, user)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) Login(ctx context.Context, email string, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("email empty")
	}

	if !helpers.IsInvalidEmail(email) {
		return Domain{}, errors.New("invalid email")
	}

	if password == "" {
		return Domain{}, errors.New("password empty")
	}

	user, err := uc.Repo.Login(ctx, email)

	if err != nil {
		return Domain{}, err
	}

	if !helpers.ComparePassword(user.HashPassword, password) {
		return Domain{}, errors.New("password wrong")
	}

	user.Token, err = uc.JWTConfig.GenerateTokenJWT(user.Id)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}