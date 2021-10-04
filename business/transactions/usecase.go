package transactions

import (
	"context"
	"errors"
)

type TransactionsUsecase struct {
	Repo Repository
}


func NewTransactionUsecase(repo Repository) Usecase {
	return &TransactionsUsecase{
		Repo: repo,
	}
}

func (t TransactionsUsecase) CreateTransactions(ctx context.Context, domain Domain) (Domain, error) {
	if domain.UserId == 0 {
		return Domain{}, errors.New("noting user")
	}

	if domain.OrderId == 0 {
		return Domain{}, errors.New("noting order id")
	}


	if domain.PaymentMethodId == 0 {
		return Domain{}, errors.New("payment method id empty")
	}

	transaction, err := t.Repo.CreateTransactions(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return transaction, nil
}


func (t TransactionsUsecase) UpdateStatusTransaction(ctx context.Context, domain Domain) (Domain, error) {
	if !(domain.Status == "Success" || domain.Status == "Failed") {
		return Domain{}, errors.New("status not valid")
	}

	if domain.Id == 0 {
		return Domain{}, errors.New("id transaction empty")
	}

	transactions, err := t.Repo.UpdateStatusTransaction(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return transactions, nil
}