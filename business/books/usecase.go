package books

import (
	"books_online_api/controllers"
	"context"
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
		return Domain{}, controllers.ID_EMPTY
	}

	resultDomain, err := b.Repo.GetOneBook(ctx, domain)

	if err != nil {
		if err.Error() == "record not found" {
			return Domain{}, controllers.RECORD_NOT_FOUND
		}
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
		return Domain{}, controllers.EMPTY_TITLE
	}

	if domain.Description == "" {
		return Domain{}, controllers.EMPTY_DESCRIPTION
	}

	resultBook, err = bookUsecase.Loc.CreateBook(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return resultBook, nil
}




