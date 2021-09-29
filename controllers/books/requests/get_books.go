package requests

import "books_online_api/business/books"

type GetBooksRequest struct {
	Keyword string
}

func (g GetBooksRequest) ToDomain() books.Domain {
	return books.Domain{
		Keyword: g.Keyword,
	}
}