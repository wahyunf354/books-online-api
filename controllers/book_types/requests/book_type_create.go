package requests

import "books_online_api/business/book_types"

type BookTypeCreate struct {
	Name string `json:"name"`
	Unit string `json:"unit"`
}

func (bt *BookTypeCreate) ToDomain()  book_types.Domain {
	return book_types.Domain{
		Name: bt.Name,
		Unit: bt.Unit,
	}
}
