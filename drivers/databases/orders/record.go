package orders

import (
	"books_online_api/business/orders"
	"books_online_api/drivers/databases/order_details"
	"gorm.io/gorm"
	"time"
)

type Orders struct {
	Id         int `gorm:"primaryKey"`
	UserId     int `gorm:"not null"`
	TotalPrice int
	Status     string

	OrderDetails []*order_details.OrderDetails `gorm:"foreignKey:OrderId;references:Id"`
	//User         *users.Users                   `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (o Orders) ToDomain(domain orders.Domain) orders.Domain {
	return orders.Domain{
		Id:          o.Id,
		UserId:      o.UserId,
		BookId:      domain.BookId,
		Qty:         domain.Qty,
		TotalPrice:  o.TotalPrice,
		StatusOrder: o.Status,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
		DeletedAt:   o.DeletedAt,
	}
}

func FromDomain(domain orders.Domain) Orders {
	return Orders{
		UserId:     domain.UserId,
		TotalPrice: domain.TotalPrice,
	}
}
