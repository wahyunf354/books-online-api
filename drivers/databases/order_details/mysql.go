package order_details

import (
	"books_online_api/business/orders"
	"books_online_api/drivers/databases/books"
	//ordersDb "books_online_api/drivers/databases/orders"
	"context"
	"gorm.io/gorm"
)

type OrderDetailRepository struct {
	Conn *gorm.DB
}

func (o OrderDetailRepository) UpdateTotalPrice(ctx context.Context, domain orders.Domain) (orders.Domain, error) {
	panic("implement me")
}

func NewMysqlOrderDetailsRepository(conn *gorm.DB) orders.OrderDetailRepository {
	return &OrderDetailRepository{Conn: conn}
}


func (o OrderDetailRepository) CreateOrder(ctx context.Context, domain orders.Domain) (orders.Domain, error) {
	var orderDetail OrderDetails
	var book books.Book

	tx := o.Conn.Begin()

	orderDetail = FromDomain(domain)

	resultBook := tx.First(&book, orderDetail.BookId)

	if resultBook.Error != nil {
		tx.Rollback()
		return orders.Domain{}, resultBook.Error
	}

	orderDetail.Price = book.Price

	resultOrderDetail := tx.Create(&orderDetail)

	if resultOrderDetail.Error != nil {
		tx.Rollback()
		return orders.Domain{}, resultOrderDetail.Error
	}

	tx.Commit()

	return orderDetail.ToDomain(domain), nil
}



