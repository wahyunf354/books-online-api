package requests

import "books_online_api/business/google_books"

type GoogleBookRequest struct {
	Keyword string
}

func (gb *GoogleBookRequest) ToDomain() google_books.Domain {
	return google_books.Domain{
		Keyword: gb.Keyword,
	}
}