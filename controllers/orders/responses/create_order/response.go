package create_order

import (
	"books_online_api/business/orders"
	"time"
)

type OrderResponse struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
	Qty int `json:"qty"`

	TotalPrice int `json:"total_price"`
	Price int `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
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