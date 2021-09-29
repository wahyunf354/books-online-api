package order_details

import (
	"books_online_api/business/orders"
	"context"
	"gorm.io/gorm"
)

type OrderDetailRepository struct {
	Conn *gorm.DB
}

func NewMysqlOrderRepository(conn *gorm.DB) orders.OrderDetailRepository {
	return &OrderDetailRepository{Conn: conn}
}


func (o OrderDetailRepository) CreateOrder(ctx context.Context, domain orders.Domain) (orders.Domain, error) {
	var orderDetail OrderDetails

	resultOrder := o.Conn.Create(&orderDetail)

	if resultOrder.Error != nil {
		return orders.Domain{}, nil
	}

	return orderDetail.ToDomain(domain), nil
}


