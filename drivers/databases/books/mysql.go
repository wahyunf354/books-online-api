package books

import (
	"books_online_api/business/books"
	"context"
	"gorm.io/gorm"
)

type BookRepository struct {
	Conn *gorm.DB
}

func (b BookRepository) CreateBook(ctx context.Context, domain books.Domain) (books.Domain, error) {
	newBook :=
}

func NewBookRepository (conn *gorm.DB) books.Repository {
	return &BookRepository{
		Conn: conn,
	}
}
