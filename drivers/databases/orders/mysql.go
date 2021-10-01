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

	// TODO: Update Total Price
	resultUpdateTotalPrice := o.Conn.First(&order, resultOrderDetail.Id)

	if resultUpdateTotalPrice.Error != nil {
		return orders.Domain{}, resultUpdateTotalPrice.Error
	}
	order.TotalPrice += domain.Price

	resultOrderSave := o.Conn.Save(&order)

	if resultOrderSave.Error != nil {
		return orders.Domain{}, resultOrderSave.Error
	}

	resultOrderDetail.TotalPrice = order.TotalPrice

	return resultOrderDetail, nil
}

func (o OrdersRepository) CheckOrderPending(ctx context.Context, domain orders.Domain) (orders.Domain, error) {

	var order Orders

	resultDb := o.Conn.Where("status = ?", "Pending").First(&order)

	if resultDb.Error != nil {
		return orders.Domain{}, resultDb.Error
	}
	return order.ToDomain(domain), nil
}
