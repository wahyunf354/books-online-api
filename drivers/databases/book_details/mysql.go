package book_details

import (
	"books_online_api/business/books"
	"context"
	"gorm.io/gorm"
)

type BookDetailsRepo struct {
	Conn *gorm.DB
}

func NewBookDetailsRepository(conn *gorm.DB) books.DetailRepository {
	return &BookDetailsRepo{
		Conn: conn,
	}
}

func (b BookDetailsRepo) CreateBook(ctx context.Context, domain books.Domain) (books.Domain, error) {
	newBookDetails := FromDomain(domain)

	resultDb := b.Conn.Create(&newBookDetails)

	if resultDb.Error != nil {
		return books.Domain{}, resultDb.Error
	}

	return newBookDetails.ToDomain(), nil
}


