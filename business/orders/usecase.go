package orders

import (
	"context"
	"time"
)

type OrderUsecase struct {
	Repo Repository
	RepoDetails OrderDetailRepository
	ContextTime time.Duration
}

func NewOrderUsecase(repository Repository, repoDetails OrderDetailRepository, timeout time.Duration) Usecase {
	return &OrderUsecase{
		Repo:        repository,
		RepoDetails: repoDetails,
		ContextTime: timeout,
	}
}

func (o OrderUsecase) CreateOrder(ctx context.Context, domain Domain) (Domain, error) {

	if domain.UserId < 0 {
		return Domain{}, nil
	}

	resultOrder, err := o.Repo.CheckOrderPanding(ctx, domain)

	if err != nil {
		order, err := o.Repo.CreateOrder(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
		return order, nil
	}

	orderDetail, err := o.RepoDetails.CreateOrder(ctx, resultOrder)

	if err != nil {
		return Domain{}, err
	}

	return orderDetail, nil
}


