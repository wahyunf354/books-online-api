package order_details

import (
	"books_online_api/business/orders"
	"books_online_api/drivers/databases/books"
	_ordersDb "books_online_api/drivers/databases/orders"
	"time"
)

type OrderDetails struct {
	Id      int        `gorm:"primaryKey"`
	OrderId int        `gorm:"not null"`
	Order 	_ordersDb.Orders `gorm:"foreignKey:OrderId"`
	BookId  int        `gorm:"not null"`
	Book    books.Book `gorm:"foreignKey:BookId"`
	Price   int        `gorm:"not null"`
	Qty     int        `gorm:"not null"`
}

func (o OrderDetails) ToDomain(domain orders.Domain) orders.Domain {
	return orders.Domain{
		Id:         o.OrderId,
		UserId:     domain.UserId,
		BookId:     o.BookId,
		Qty:        o.Qty,
		TotalPrice: domain.TotalPrice + o.Price,
		Price:      o.Price,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
		DeletedAt:  nil,
	}
}
