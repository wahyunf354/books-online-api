package payment_methods

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	Id int
	Name string
	Fee int
	Address string
	Author string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	CreatePaymentMethod(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	CreatePaymentMethod(ctx context.Context, domain Domain) (Domain, error)
}
