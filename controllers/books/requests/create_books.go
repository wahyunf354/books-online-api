package requests

import (
	"books_online_api/business/books"
	"mime/multipart"
)

type BookRequest struct {
	Title       string                  `json:"title" form:"title"`
	BookTypeId  int                     `json:"book_type_id" form:"book_type_id"`
	Price       int                     `json:"price" form:"price"`
	UserId      int                     `json:"user_id" form:"user_id"`
	Description string                  `json:"description" form:"description"`
	PageCount   int                     `json:"page_count" form:"page_count"`
	FileBook    *multipart.FileHeader   `json:"file_book" form:"file_book"`
	ImagesBook  []*multipart.FileHeader `json:"images" form:"images"`
}

func (cb *BookRequest) ToDomain() books.Domain {
	return books.Domain{
		Title:       cb.Title,
		BookTypeId:  cb.BookTypeId,
		Price:       cb.Price,
		UserId:      cb.UserId,
		Description: cb.Description,
		PageCount:   cb.PageCount,
		FileBook:    cb.FileBook,
		FileCover:   cb.ImagesBook,
	}
}

func FromDomain(domain books.Domain) BookRequest {
	return BookRequest{
		Title:       domain.Title,
		BookTypeId:  domain.BookTypeId,
		Price:       domain.Price,
		UserId:      domain.UserId,
		Description: domain.Description,
		PageCount:   domain.PageCount,
		FileBook:    domain.FileBook,
	}
}
