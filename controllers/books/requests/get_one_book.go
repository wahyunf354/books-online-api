package requests

import (
	"books_online_api/business/books"
)

type GetOneBookRequest struct {
	Id int
}

func (g GetOneBookRequest) ToDomain() books.Domain {
	return books.Domain{
		Id: g.Id,
	}
}