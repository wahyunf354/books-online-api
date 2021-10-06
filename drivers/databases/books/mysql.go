package books

import (
	"books_online_api/business/books"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepository struct {
	Conn *gorm.DB
	BookDetailsRepository books.DetailRepository
}

func NewBookRepository (conn *gorm.DB, bookDetailRepo books.DetailRepository) books.Repository {
	return &BookRepository{
		Conn: conn,
		BookDetailsRepository: bookDetailRepo,
	}
}

func (b BookRepository) GetOneBook(ctx context.Context, domain books.Domain) (books.Domain, error) {

	var bookResult Book

	resultDb := b.Conn.Preload(clause.Associations).First(&bookResult, domain.Id)

	if resultDb.Error != nil {
		return books.Domain{}, resultDb.Error
	}

	return bookResult.ToDomain(domain), nil
}

func (b BookRepository) GetBooks(ctx context.Context, domain books.Domain) ([]books.Domain, error) {

	var booksResult []Book

	result := b.Conn.Preload(clause.Associations).Find(&booksResult)

	if result.Error != nil {
		return []books.Domain{}, result.Error
	}

	return ToListDomain(booksResult, domain), nil
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


