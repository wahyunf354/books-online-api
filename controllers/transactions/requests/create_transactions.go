package requests

import (
	"books_online_api/business/transactions"
)

type CreateTransactionRequest struct {
	OrderId         int    `json:"order_id"`
	PaymentMethodId int    `json:"payment_method_id"`
	UserId          int    `json:"user_id"`
}

func (c CreateTransactionRequest) ToDomain() transactions.Domain {
	return transactions.Domain{
		OrderId:         c.OrderId,
		PaymentMethodId: c.PaymentMethodId,
		UserId:          c.UserId,
	}
}

func FromDomain(domain transactions.Domain) CreateTransactionRequest {
	return CreateTransactionRequest{
		OrderId:         domain.OrderId,
		PaymentMethodId: domain.PaymentMethodId,
		UserId:          domain.UserId,
	}
}
