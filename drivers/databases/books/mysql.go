package books

import (
	"books_online_api/business/books"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type BookRepository struct {
	Conn *gorm.DB
	BookDetailsRepository books.DetailRepository
}

func (b BookRepository) GetBooks(ctx context.Context) (books.Domain, error) {

	fmt.Println("Hai")
	var booksResult Book

	result := b.Conn.Find(&booksResult)

	if result.Error != nil {
		return books.Domain{}, result.Error
	}

	fmt.Println(booksResult)

	return books.Domain{}, nil
}

func (b BookRepository) CreateBook(ctx context.Context, domain books.Domain) (books.Domain, error) {
	newBook := FromDomain(domain)

	resultBook := b.Conn.Create(&newBook)

	if resultBook.Error != nil {
		return  books.Domain{}, resultBook.Error
	}

	resultDetail, err := b.BookDetailsRepository.CreateBook(ctx, newBook.ToDomain(domain))
	if err != nil {
		return books.Domain{}, err
	}
	return resultDetail, nil
}

func NewBookRepository (conn *gorm.DB, bookDetailRepo books.DetailRepository) books.Repository {
	return &BookRepository{
		Conn: conn,
		BookDetailsRepository: bookDetailRepo,
	}
}
