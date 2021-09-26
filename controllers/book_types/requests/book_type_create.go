package requests

import "books_online_api/business/book_types"

type BookTypeCreate struct {
	Name string `json:"name"`
	Unit string `json:"unit"`
}

func (bt *BookTypeCreate) ToDomain()  book_types.BookTypeDomain {
	return book_types.BookTypeDomain{
		Name: bt.Name,
		Unit: bt.Unit,
	}
}
