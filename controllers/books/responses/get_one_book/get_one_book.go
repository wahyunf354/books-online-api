package get_one_book

import (
	"books_online_api/business/books"
	"time"
)

type OneBookResponse struct {
	Id int
	Title string
	UserId int
	Price int

	BookType interface{}
	BookDetail interface{}
	ImageBooks interface{}

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt interface{}
}

func FromDomain(domain books.Domain) OneBookResponse {
	return OneBookResponse{
		Id:         domain.Id,
		Title:      domain.Title,
		UserId:     domain.UserId,
		Price:      domain.Price,
		BookType:   domain.BookType,
		BookDetail: domain.BookDetail,
		ImageBooks: domain.ImageBooks,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
		DeletedAt:  domain.DeletedAt,
	}
}