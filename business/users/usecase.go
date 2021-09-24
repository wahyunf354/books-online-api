package users

import (
	"books_online_api/helpers"
	"context"
	"errors"
	"time"
)

type UserUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUseCase(repo Repository, timeout time.Duration) Usecase {
	return &UserUseCase{
		Repo:           repo,
		contextTimeout: timeout,
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
