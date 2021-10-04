package transactions

import (
	"books_online_api/business/orders"
	"books_online_api/business/payment_methods"
	"books_online_api/business/transactions"
	"context"
	"gorm.io/gorm"
)

type TransactionsRepository struct {
	Conn *gorm.DB
	PaymentMethodRepo payment_methods.Repository
	OrdersRepo orders.Repository
}

func NewTransactionsRepository(conn *gorm.DB, repoPaymentMethod payment_methods.Repository, repoOrder orders.Repository) transactions.Repository {
	return TransactionsRepository{Conn: conn, PaymentMethodRepo: repoPaymentMethod, OrdersRepo: repoOrder}
}

func (t TransactionsRepository) CreateTransactions(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	var newTransaction Transactions
	var paymentMethod payment_methods.Domain
	var order orders.Domain
	var err error

	paymentMethod.Id = domain.PaymentMethodId
	order.Id = domain.OrderId

	paymentMethod, err = t.PaymentMethodRepo.GetOnePaymentMethod(ctx, paymentMethod)

	if err != nil {
		return transactions.Domain{}, err
	}

	order, err = t.OrdersRepo.GetOneOrder(ctx, order)

	if err != nil {
		return transactions.Domain{}, err
	}

	domain.TotalPrice += order.TotalPrice
	domain.TotalPrice += paymentMethod.Fee

	newTransaction = FromDomain(domain)

	result := t.Conn.Create(&newTransaction)

	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	return newTransaction.ToDomain(), nil
}

