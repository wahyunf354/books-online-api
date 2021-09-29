package books

import (
	"context"
	"errors"
	"time"
)

type BookUsecase struct {
	Loc Localy
	Repo Repository
	ContextTimeout time.Duration
}

func NewBookUsecase (loc Localy, repo Repository,  timeout time.Duration) Usecase {
	return &BookUsecase{
		Loc: loc,
		Repo: repo,
		ContextTimeout: timeout,
	}
}

func (b *BookUsecase) GetOneBook(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Id == 0 {
		return Domain{}, errors.New("Id empty")
	}

	resultDomain, err := b.Repo.GetOneBook(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return resultDomain, nil
}

func (b *BookUsecase) GetBooks(ctx context.Context, domain Domain) ([]Domain, error) {

	resultRepo,err := b.Repo.GetBooks(ctx, domain)

	if err != nil {
		return []Domain{}, err
	}

	return resultRepo, nil
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




