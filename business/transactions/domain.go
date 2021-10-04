package transactions

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	Id int
	OrderId int
	PaymentMethodId int
	Status string
	TotalPrice int
	UserId int
	
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	CreateTransactions(ctx context.Context, domain Domain) (Domain, error)
	UpdateStatusTransaction(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	CreateTransactions(ctx context.Context, domain Domain) (Domain, error)
	UpdateStatusTransaction(ctx context.Context, domain Domain) (Domain, error)
}


















