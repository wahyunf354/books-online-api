package orders

import (
	"books_online_api/controllers"
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

	if domain.UserId < 1 {
		return Domain{}, controllers.FORBIDDEN_USER
	}

	if domain.BookId == 0 {
		return Domain{}, controllers.EMPTY_BOOK_ID
	}

	if domain.Qty == 0 {
		return Domain{}, controllers.EMPTY_QTY
	}

	resultOrder, err := o.Repo.CheckOrderPending(ctx, domain)

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

	result, err := o.Repo.UpdateTotalPrice(ctx, orderDetail)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}


