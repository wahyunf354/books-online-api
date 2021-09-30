package orders

import (
	"books_online_api/business/orders"
	"context"
	"gorm.io/gorm"
)

type OrdersRepository struct {
	Conn *gorm.DB
	RepoOrderDetail orders.OrderDetailRepository
}

func NewMysqlOrderRepository(conn *gorm.DB, repo orders.OrderDetailRepository) orders.Repository {
	return &OrdersRepository{Conn: conn, RepoOrderDetail: repo}
}


func (o OrdersRepository) CreateOrder(ctx context.Context, domain orders.Domain) (orders.Domain, error) {
	var order Orders

	order = FromDomain(domain)

	order.Status = "Pending"

	resultOrder := o.Conn.Create(&order)

	if resultOrder.Error != nil {
		return orders.Domain{}, resultOrder.Error
	}

	resultOrderDetail, err := o.RepoOrderDetail.CreateOrder(ctx, order.ToDomain(domain))

	if err != nil {
		return orders.Domain{}, err
	}

	return resultOrderDetail, nil
}

func (o OrdersRepository) CheckOrderPanding(ctx context.Context, domain orders.Domain) (orders.Domain, error) {

	var order Orders

	resultDb := o.Conn.Where("status = ?", "Pending").First(&order)

	if resultDb.Error != nil {
		return orders.Domain{}, resultDb.Error
	}
	return order.ToDomain(domain), nil
}
