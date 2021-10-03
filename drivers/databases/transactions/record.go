package transactions

import (
	"books_online_api/business/transactions"
	"gorm.io/gorm"
	"time"
)

type Transactions struct {
	Id              int `gorm:"primaryKey"`
	OrderId         int
	PaymentMethodId int
	Status          string `gorm:"default=Pending"`
	TotalPrice      int
	UserId          int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func FromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		OrderId:         domain.OrderId,
		PaymentMethodId: domain.PaymentMethodId,
		Status:          domain.Status,
		TotalPrice:      domain.TotalPrice,
		UserId:          domain.UserId,
	}
}

func (t Transactions) ToDomain() transactions.Domain {
	return transactions.Domain{
		Id:              t.Id,
		OrderId:         t.OrderId,
		PaymentMethodId: t.PaymentMethodId,
		Status:          t.Status,
		TotalPrice:      t.TotalPrice,
		UserId:          t.UserId,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
		DeletedAt:       t.DeletedAt,
	}
}
