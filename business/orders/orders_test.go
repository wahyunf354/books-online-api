package orders_test

import (
	"books_online_api/business/orders"
	_mocksOrderRepository "books_online_api/business/orders/mocks"
	"books_online_api/controllers"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var orderRepository _mocksOrderRepository.Repository
var orderDetailsRepository _mocksOrderRepository.OrderDetailRepository

var orderService orders.Usecase
var orderDomain orders.Domain

func setup() {
	orderService = orders.NewOrderUsecase(&orderRepository, &orderDetailsRepository, time.Hour*1)
	orderDomain = orders.Domain{
		Id:          1,
		UserId:      4,
		BookId:      4,
		Qty:         2,
		TotalPrice:  70000,
		Price:       25000,
		StatusOrder: "Pending",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   nil,
	}
}

func TestOrderUsecase_CreateOrder(t *testing.T) {
	t.Run("Test 1 | valid create order", func(t *testing.T) {
		setup()
		orderRepository.On("CheckOrderPending", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderRepository.On("CreateOrder", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderDetailsRepository.On("CreateOrder", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderRepository.On("UpdateTotalPrice", mock.Anything, mock.Anything).Return(orderDomain, nil)
		_, err := orderService.CreateOrder(context.Background(), orders.Domain{UserId: 2, BookId: 2, Qty: 4})
		assert.Nil(t, err)
	})
	t.Run("Test 2 | book id empty", func(t *testing.T) {
		setup()
		orderRepository.On("CheckOrderPending", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderRepository.On("CreateOrder", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderDetailsRepository.On("CreateOrder", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderRepository.On("UpdateTotalPrice", mock.Anything, mock.Anything).Return(orderDomain, nil)
		_, err := orderService.CreateOrder(context.Background(), orders.Domain{UserId: 2, BookId: 0, Qty: 4})
		assert.NotNil(t, err)
		assert.Equal(t, controllers.EMPTY_BOOK_ID, err)
	})
	t.Run("Test 3 | qty empty", func(t *testing.T) {
		setup()
		orderRepository.On("CheckOrderPending", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderRepository.On("CreateOrder", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderDetailsRepository.On("CreateOrder", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderRepository.On("UpdateTotalPrice", mock.Anything, mock.Anything).Return(orderDomain, nil)
		_, err := orderService.CreateOrder(context.Background(), orders.Domain{UserId: 2, BookId: 2, Qty: 0})
		assert.NotNil(t, err)
		assert.Equal(t, controllers.EMPTY_QTY, err)
	})

	t.Run("Test 4 | user empty", func(t *testing.T) {
		setup()
		orderRepository.On("CheckOrderPending", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderRepository.On("CreateOrder", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderDetailsRepository.On("CreateOrder", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderRepository.On("UpdateTotalPrice", mock.Anything, mock.Anything).Return(orderDomain, nil)
		_, err := orderService.CreateOrder(context.Background(), orders.Domain{UserId: 0, BookId: 2, Qty: 3})
		assert.NotNil(t, err)
		assert.Equal(t, controllers.FORBIDDEN_USER, err)
	})

	t.Run("Test 6 | valid add order detail tanpa order id", func(t *testing.T) {
		setup()
		orderRepository.On("CheckOrderPending", mock.Anything, mock.Anything).Return(orders.Domain{}, controllers.RECORD_NOT_FOUND)
		orderRepository.On("CreateOrder", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderDetailsRepository.On("CreateOrder", mock.Anything, mock.Anything).Return(orderDomain, nil)
		orderRepository.On("UpdateTotalPrice", mock.Anything, mock.Anything).Return(orderDomain, nil)
		_, err := orderService.CreateOrder(context.Background(), orders.Domain{UserId: 1, BookId: 2, Qty: 3})
		assert.Nil(t, err)
	})

}
