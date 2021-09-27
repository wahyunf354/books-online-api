package books

import (
	"books_online_api/business/books"
	"context"
	"gorm.io/gorm"
)

type BookRepository struct {
	Conn *gorm.DB
	BookDetailsRepository books.DetailRepository
}

func (b BookRepository) CreateBook(ctx context.Context, domain books.Domain) (books.Domain, error) {
	newBook := FromDomain(domain)

	resultBook := b.Conn.Create(&newBook)

	if resultBook.Error != nil {
		return  books.Domain{}, resultBook.Error
	}

	resultDetail, err := b.BookDetailsRepository.CreateBook(ctx, newBook.ToDomain())
	if err != nil {
		return books.Domain{}, err
	}
	return resultDetail, nil
}

func NewBookRepository (conn *gorm.DB) books.Repository {
	return &BookRepository{
		Conn: conn,
	}
}
