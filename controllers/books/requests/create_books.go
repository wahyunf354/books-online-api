package requests

import (
	"books_online_api/business/books"
	"mime/multipart"
)

type BookRequest struct {
	Title string `json:"title"`
	BookTypeId int `json:"book_type_id"`
	Price int `json:"price"`
	UserId int `json:"user_id"`
	Description string `json:"description"`
	PageCount int `json:"page_count"`
	FileBook multipart.FileHeader
}

func (cb *BookRequest) ToDomain() books.Domain {
	return books.Domain{
		Title: cb.Title,
		BookTypeId: cb.BookTypeId,
		Price: cb.Price,
		UserId: cb.UserId,
		Description: cb.Description,
		PageCount: cb.PageCount,
		FileBook: cb.FileBook,
	}
}

func FromDomain(domain books.Domain) BookRequest {
	return BookRequest{
		Title: domain.Title,
		BookTypeId: domain.BookTypeId,
		Price: domain.Price,
		UserId: domain.UserId,
		Description: domain.Description,
		PageCount: domain.PageCount,
		FileBook: domain.FileBook,
	}
}