package orders

import (
	"books_online_api/business/orders"
	"books_online_api/controllers"
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

	result, err := o.UpdateTotalPrice(ctx, resultOrderDetail)

	if err != nil {
		return orders.Domain{}, err
	}

	return result, nil
}

func (o OrdersRepository) UpdateTotalPrice(ctx context.Context, domain orders.Domain) (orders.Domain, error) {
	var order Orders
	resultUpdateTotalPrice := o.Conn.First(&order, domain.Id)

	if resultUpdateTotalPrice.Error != nil {
		return orders.Domain{}, resultUpdateTotalPrice.Error
	}
	order.TotalPrice = order.TotalPrice + domain.Price

	resultOrderSave := o.Conn.Save(&order)

	if resultOrderSave.Error != nil {
		return orders.Domain{}, resultOrderSave.Error
	}

	domain.TotalPrice = order.TotalPrice
	return domain, nil
}

func (o OrdersRepository) CheckOrderPending(ctx context.Context, domain orders.Domain) (orders.Domain, error) {

	var order Orders

	resultDb := o.Conn.Where("status = ?", "Pending").First(&order)

	if resultDb.Error != nil {
		return orders.Domain{}, controllers.RECORD_NOT_FOUND
	}
	return order.ToDomain(domain), nil
}


func (o OrdersRepository) GetOneOrder(ctx context.Context, domain orders.Domain) (orders.Domain, error) {
	var ordersResult Orders

	result := o.Conn.First(&ordersResult, domain.Id)

	if result.Error != nil {
		return orders.Domain{}, result.Error
	}

	return ordersResult.ToDomain(domain), nil
}