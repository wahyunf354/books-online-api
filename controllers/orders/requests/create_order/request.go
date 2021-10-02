package create_order

import "books_online_api/business/orders"

type OrderRequest struct {
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
	Qty    int `json:"qty"`
}

func (o OrderRequest) ToDomain() orders.Domain {
	return orders.Domain{
		UserId: o.UserId,
		BookId: o.BookId,
		Qty:    o.Qty,
	}
}
