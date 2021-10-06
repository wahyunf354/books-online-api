package book_types

import (
	"books_online_api/controllers"
	"context"
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
		return Domain{}, controllers.NAME_BOOK_TYPE_EMPTY
	}

	if bookType.Unit == "" {
		return Domain{}, controllers.UNIT_BOOK_TYPE_EMPTY
	}

	bookType, err := uc.Repo.CreateBookType(ctx, bookType)

	if err != nil {
		return Domain{}, err
	}

	return bookType, nil
}