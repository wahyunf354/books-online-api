package order_details

import (
	"books_online_api/business/orders"
	"time"
)

type OrderDetails struct {
	Id        int              `gorm:"primaryKey"`
	OrderId   int              `gorm:"not null"`
	//Order     _ordersDb.Orders `gorm:"foreignKey:OrderId"`
	BookId    int              `gorm:"not null"`
	Price     int              `gorm:"not null"`
	Qty       int              `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (o OrderDetails) ToDomain(domain orders.Domain) orders.Domain {
	return orders.Domain{
		Id:         o.OrderId,
		UserId:     domain.UserId,
		BookId:     o.BookId,
		Qty:        o.Qty,
		TotalPrice: domain.TotalPrice + o.Price,
		Price:      o.Price,
		CreatedAt:  o.CreatedAt,
		UpdatedAt:  o.UpdatedAt,
		DeletedAt:  o.DeletedAt,
	}
}
