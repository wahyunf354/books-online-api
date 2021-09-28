package book_types

import (
	"context"
	"errors"
	"time"
)

type BookTypesUseCase struct {
	Repo Repository
	ContextTimeout time.Duration
}

func NewBooksUseCase(repo Repository, timeout time.Duration) Usecase {
	return &BookTypesUseCase{
		Repo: repo,
		ContextTimeout: timeout,
	}
}

func (uc *BookTypesUseCase) CreateBookType(ctx context.Context, bookType Domain) (Domain, error) {
	if bookType.Name == "" {
		return Domain{}, errors.New("name book type empty")
	}

	if bookType.Unit == "" {
		return Domain{}, errors.New("unit book type empty")
	}

	bookType, err := uc.Repo.CreateBookType(ctx, bookType)

	if err != nil {
		return Domain{}, err
	}

	return bookType, nil
}