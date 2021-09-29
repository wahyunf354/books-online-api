package get_books

import (
	"books_online_api/business/books"
	"time"
)

type BookResponse struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	UserId       int    `json:"user_id"`
	Description  string `json:"description"`
	UrlBook      string `json:"url_book"`
	Price        int    `json:"price"`
	BookTypeId   int    `json:"book_type_id"`
	PageCount    int    `json:"page_count"`
	BookDetailId int    `json:"book_detail_id"`

	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
}

func FromListDomain(domains []books.Domain) []BookResponse {
	var result []BookResponse
	for _,domain := range domains{
		result = append(result, BookResponse{
			Id:           domain.Id,
			Title:        domain.Title,
			UserId:       domain.UserId,
			Description:  domain.Description,
			UrlBook:      domain.UrlBook,
			Price:        domain.Price,
			BookTypeId:   domain.BookTypeId,
			PageCount:    domain.PageCount,
			BookDetailId: domain.BookDetailId,
			CreatedAt:    domain.CreatedAt,
			UpdatedAt:    domain.UpdatedAt,
			DeletedAt:    domain.DeletedAt,
		})
	}
	return result
}