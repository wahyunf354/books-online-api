package payment_methods

import (
	"context"
	"errors"
)

type PaymentMethodUsecase struct {
	Repo Repository
}

func NewPaymentMethodUsecase(repo Repository) Usecase {
	return &PaymentMethodUsecase{Repo: repo}
}

func (p PaymentMethodUsecase) CreatePaymentMethod(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("Name is empty")
	}
	if domain.Address == "" {
		return Domain{}, errors.New("Address is empty")
	}
	if domain.Author == "" {
		return Domain{}, errors.New("Author is empty")
	}

	resultDomain, err := p.Repo.CreatePaymentMethod(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return resultDomain, nil
}
