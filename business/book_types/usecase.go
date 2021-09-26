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

func (uc *BookTypesUseCase) CreateBookType(ctx context.Context, bookType BookTypeDomain) (BookTypeDomain, error) {
	if bookType.Name == "" {
		return BookTypeDomain{}, errors.New("name book type empty")
	}

	if bookType.Unit == "" {
		return BookTypeDomain{}, errors.New("unit book type empty")
	}

	bookType, err := uc.Repo.CreateBookType(ctx, bookType)

	if err != nil {
		return BookTypeDomain{}, err
	}

	return bookType, nil
}