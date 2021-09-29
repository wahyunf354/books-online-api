package orders

import (
	"context"
	"time"
)

type Domain struct {
	Id int
	UserId int
	BookId int
	Qty int

	TotalPrice int
	Price int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt interface{}
}

type Usecase interface {
	CreateOrder(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	CreateOrder(ctx context.Context, domain Domain) (Domain, error)
}

type OrderDetailRepository interface {
	CreateOrder(ctx context.Context, domain Domain) (Domain, error)
}