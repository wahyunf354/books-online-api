package orders

import (
	"context"
	"time"
)

type OrderUsecase struct {
	Repo Repository
	ContextTime time.Duration
}

func NewOrderUsecase(repository Repository, timeout time.Duration) Usecase {
	return &OrderUsecase{
		Repo:        repository,
		ContextTime: timeout,
	}
}

func (o OrderUsecase) CreateOrder(ctx context.Context, domain Domain) (Domain, error) {

	if domain.UserId < 1 {
		return Domain{}, nil
	}

	order, err := o.Repo.CreateOrder(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return order, nil
}


