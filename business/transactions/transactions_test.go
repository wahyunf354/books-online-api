package transactions_test

import (
	"books_online_api/business/transactions"
	_mocksTransactions "books_online_api/business/transactions/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"time"
)

var transactionsRepository _mocksTransactions.Repository

var transactionsService transactions.Usecase
var transactionsDomain transactions.Domain

func setup() {
	transactionsService = transactions.NewTransactionUsecase(&transactionsRepository)
	transactionsDomain = transactions.Domain{
		Id:              1,
		OrderId:         1,
		PaymentMethodId: 2,
		Status:          "Pending",
		TotalPrice:      25000,
		UserId:          3,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       gorm.DeletedAt{},
	}
}

func TestTransactionsUsecase_CreateTransactions(t *testing.T) {
	t.Run("Test case 1 | success create transaction", func(t *testing.T) {
		setup()
		transactionsRepository.On("CreateTransactions", mock.Anything, mock.Anything).Return(transactionsDomain, nil)
		_, err := transactionsService.CreateTransactions(context.Background(), transactions.Domain{
			OrderId:         1,
			PaymentMethodId: 2,
			TotalPrice:      25000,
			UserId:          3,
		})
		assert.Nil(t, err)
	})
}

func setupUpdateTransactions() {
	transactionsService = transactions.NewTransactionUsecase(&transactionsRepository)
	transactionsDomain = transactions.Domain{
		Id:              1,
		OrderId:         1,
		PaymentMethodId: 2,
		Status:          "Success",
		TotalPrice:      25000,
		UserId:          3,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       gorm.DeletedAt{},
	}
}

func TestTransactionsUsecase_UpdateStatusTransaction(t *testing.T) {
	t.Run("Test case 1 | success update transactions", func(t *testing.T) {
		setupUpdateTransactions()
		transactionsRepository.On("UpdateStatusTransaction", mock.Anything, mock.Anything).Return(transactionsDomain, nil)
		transactions, err := transactionsService.UpdateStatusTransaction(context.Background(), transactions.Domain{
			Id:              1,
			Status:          "Success",
		})
		assert.Nil(t, err)
		assert.Equal(t, "Success", transactions.Status)
	})
}