package books

import (
	"context"
	"errors"
	"time"
)

type BookUsecase struct {
	Repo Repository
	Loc Localy
	ContextTimeout time.Duration
}

func NewBookUsecase (repo Repository, loc Localy, timeout time.Duration) Usecase {
	return &BookUsecase{
		Repo: repo,
		Loc: loc,
		ContextTimeout: timeout,
	}
}

func (bookUsecase *BookUsecase) CreateBook(ctx context.Context, book Domain) (Domain, error) {

	var resultBook Domain
	var err error

	if book.Title == "" {
		return Domain{}, errors.New("Title empty")
	}

	if book.Description == "" {
		return Domain{}, errors.New("Description empty")
	}

	if book.UserId == 0 {
		return Domain{}, errors.New("User id empty")
	}

	resultBook, err = bookUsecase.Loc.CreateBook(ctx, book)

	if err != nil {
		return Domain{}, err
	}

	return resultBook, nil
}




