package create_order

import (
	"books_online_api/business/orders"
	"time"
)

type OrderResponse struct {
	Id int
	UserId int
	BookId int
	Qty int

	TotalPrice int
	Price int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt interface{}
}

func FromDomain(domain orders.Domain) OrderResponse {
	return OrderResponse{
		Id:         domain.Id,
		UserId:     domain.UserId,
		BookId:     domain.BookId,
		Qty:        domain.Qty,
		TotalPrice: domain.TotalPrice,
		Price:      domain.Price,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
	}
}