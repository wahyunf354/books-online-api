package requests

import (
	"books_online_api/business/transactions"
)

type UpdateStatusTransactionRequest struct {
	Status        string `json:"status"`
	TransactionId int    `json:"transaction_id"`
}

func (u UpdateStatusTransactionRequest) ToDomain() transactions.Domain {
	return transactions.Domain{
		Id:     u.TransactionId,
		Status: u.Status,
	}
}
