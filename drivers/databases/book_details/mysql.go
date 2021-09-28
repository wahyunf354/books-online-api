package book_details

import (
	"books_online_api/business/books"
	"context"
	"gorm.io/gorm"
)

type BookDetailsRepo struct {
	Conn *gorm.DB
	ImageBooksLocal books.ImageBooksLocaly
}

func NewBookDetailsRepository(conn *gorm.DB, imageBooksLocal books.ImageBooksLocaly) books.DetailRepository {
	return &BookDetailsRepo{
		Conn: conn,
		ImageBooksLocal: imageBooksLocal,
	}
}

func (b BookDetailsRepo) CreateBook(ctx context.Context, domain books.Domain) (books.Domain, error) {
	newBookDetails := FromDomain(domain)

	resultDb := b.Conn.Create(&newBookDetails)

	if resultDb.Error != nil {
		return books.Domain{}, resultDb.Error
	}

	result, err := b.ImageBooksLocal.UploadImages(ctx, newBookDetails.ToDomain(domain))

	if err != nil {
		return books.Domain{}, err
	}

	return result, nil
}


