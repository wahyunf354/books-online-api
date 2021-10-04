package update_transaction

import (
	"books_online_api/business/transactions"
	"gorm.io/gorm"
	"time"
)

type UpdateStatusTransaction struct {
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

func FromDomain(domain transactions.Domain) UpdateStatusTransaction {
	return UpdateStatusTransaction{
		Id:              domain.Id,
		OrderId:         domain.OrderId,
		PaymentMethodId: domain.PaymentMethodId,
		Status:          domain.Status,
		TotalPrice:      domain.TotalPrice,
		UserId:          domain.UserId,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
		DeletedAt:       domain.DeletedAt,
	}
}
