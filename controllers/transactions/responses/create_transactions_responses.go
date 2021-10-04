package responses

import (
	"books_online_api/business/transactions"
	"gorm.io/gorm"
	"time"
)

type CreateTransactionsResponses struct {
	Id int `json:"id"`
	OrderId int `json:"order_id"`
	PaymentMethodId int `json:"payment_method_id"`
	Status string `json:"status"`
	TotalPrice int `json:"total_price"`
	UserId int `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain transactions.Domain) CreateTransactionsResponses {
	return CreateTransactionsResponses{
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