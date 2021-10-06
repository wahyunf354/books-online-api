package payment_methods_test

import (
	"books_online_api/business/payment_methods"
	_mocksPaymentMethod "books_online_api/business/payment_methods/mocks"
	"books_online_api/controllers"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"time"
)

var paymentMethodRepository _mocksPaymentMethod.Repository

var paymentMethodService payment_methods.Usecase
var paymentMethodDomain payment_methods.Domain

func setup(){
	paymentMethodService = payment_methods.NewPaymentMethodUsecase(&paymentMethodRepository)
	paymentMethodDomain = payment_methods.Domain{
		Id:        1,
		Name:      "BNI",
		Fee:       0,
		Address:   "34235234",
		Author:    "Wahyu",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}

func TestPaymentMethodUsecase_CreatePaymentMethod(t *testing.T) {
	t.Run("Test Case 1 | create payment method", func(t *testing.T) {
		setup()
		paymentMethodRepository.On("CreatePaymentMethod", mock.Anything, mock.Anything).Return(paymentMethodDomain, nil)
		paymentMethod, err := paymentMethodService.CreatePaymentMethod(context.Background(), payment_methods.Domain{
			Name:      "BRI",
			Fee:       123,
			Address:   "2917834",
			Author:    "Wahyu",
		})
		assert.Nil(t, err)
		assert.Equal(t, 1, paymentMethod.Id)
	})

	t.Run("Test Case 2 | Name empty", func(t *testing.T) {
		setup()
		paymentMethodRepository.On("CreatePaymentMethod", mock.Anything, mock.Anything).Return(paymentMethodDomain, nil)
		paymentMethod, err := paymentMethodService.CreatePaymentMethod(context.Background(), payment_methods.Domain{
			Name:      "",
			Fee:       123,
			Address:   "2917834",
			Author:    "Wahyu",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controllers.NAME_EMPTY, err)
		assert.Equal(t, "", paymentMethod.Name)
	})

	t.Run("Test Case 3 | Address empty", func(t *testing.T) {
		setup()
		paymentMethodRepository.On("CreatePaymentMethod", mock.Anything, mock.Anything).Return(paymentMethodDomain, nil)
		paymentMethod, err := paymentMethodService.CreatePaymentMethod(context.Background(), payment_methods.Domain{
			Name:      "BRI",
			Fee:       123,
			Address:   "",
			Author:    "Wahyu",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controllers.ADDRESS_EMPTY, err)
		assert.Equal(t, "", paymentMethod.Address)
	})

	t.Run("Test Case 4 | Author empty", func(t *testing.T) {
		setup()
		paymentMethodRepository.On("CreatePaymentMethod", mock.Anything, mock.Anything).Return(paymentMethodDomain, nil)
		paymentMethod, err := paymentMethodService.CreatePaymentMethod(context.Background(), payment_methods.Domain{
			Name:      "BRI",
			Fee:       123,
			Address:   "134335423",
			Author:    "",
		})
		assert.NotNil(t, err)
		assert.Equal(t, controllers.AUTHOR_EMPTY, err)
		assert.Equal(t, "", paymentMethod.Author)
	})

}