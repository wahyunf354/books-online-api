package payment_methods

import (
	"books_online_api/controllers"
	"context"
)

type PaymentMethodUsecase struct {
	Repo Repository
}

func NewPaymentMethodUsecase(repo Repository) Usecase {
	return &PaymentMethodUsecase{Repo: repo}
}

func (p PaymentMethodUsecase) CreatePaymentMethod(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, controllers.NAME_EMPTY
	}
	if domain.Address == "" {
		return Domain{}, controllers.ADDRESS_EMPTY
	}
	if domain.Author == "" {
		return Domain{}, controllers.AUTHOR_EMPTY
	}

	resultDomain, err := p.Repo.CreatePaymentMethod(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return resultDomain, nil
}
