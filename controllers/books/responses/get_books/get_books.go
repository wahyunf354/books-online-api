package get_books

import (
	"books_online_api/business/books"
	"time"
)

type BookResponse struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	UserId      int         `json:"user_id"`
	Price       int         `json:"price"`
	BookDetails interface{} `json:"book_details"`
	BookType    interface{} `json:"book_type"`
	ImageBooks  interface{} `json:"images_book"`

	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
}

func FromListDomain(domains []books.Domain) []BookResponse {
	var result []BookResponse
	for _, domain := range domains {
		result = append(result, BookResponse{
			Id:          domain.Id,
			Title:       domain.Title,
			UserId:      domain.UserId,
			Price:       domain.Price,
			BookDetails: domain.BookDetail,
			BookType:    domain.BookType,
			ImageBooks:  domain.ImageBooks,

			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
			DeletedAt: domain.DeletedAt,
		})
	}
	return result
}
