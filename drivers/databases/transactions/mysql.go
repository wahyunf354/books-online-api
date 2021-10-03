package transactions

import (
	"books_online_api/business/transactions"
	"context"
	"gorm.io/gorm"
)

type TransactionsRepository struct {
	Conn *gorm.DB
}

func NewTransactionsRepository(conn *gorm.DB) transactions.Repository {
	return TransactionsRepository{Conn: conn}
}


func (t TransactionsRepository) CreateTransactions(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	var newTransaction Transactions
	newTransaction = FromDomain(domain)

	result := t.Conn.Create(&newTransaction)

	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	return newTransaction.ToDomain(), nil
}

