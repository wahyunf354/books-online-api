package books

import (
	"context"
	"errors"
	"time"
)

type BookUsecase struct {
	Loc Localy
	ContextTimeout time.Duration
}

func NewBookUsecase (loc Localy, timeout time.Duration) Usecase {
	return &BookUsecase{
		Loc: loc,
		ContextTimeout: timeout,
	}
}

func (bookUsecase *BookUsecase) CreateBook(ctx context.Context, domain Domain) (Domain, error) {

	var resultBook Domain
	var err error

	if domain.Title == "" {
		return Domain{}, errors.New("Title empty")
	}

	if domain.Description == "" {
		return Domain{}, errors.New("Description empty")
	}

	if domain.UserId == 0 {
		return Domain{}, errors.New("User id empty")
	}

	if len(domain.FileCover) < 1 {
		return Domain{}, errors.New("image cover empty")
	}

	resultBook, err = bookUsecase.Loc.CreateBook(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return resultBook, nil
}




