package image_books

import (
	"books_online_api/business/books"
	"context"
	"gorm.io/gorm"
)

type ImageBooksRepository struct {
	Conn *gorm.DB
}

func (i ImageBooksRepository) UploadImages(ctx context.Context, domain books.Domain) (books.Domain, error) {
	newImageBooks := FromDomain(domain)

	for _, imageBook := range newImageBooks {
		result := i.Conn.Create(&imageBook)
		if result.Error != nil {
			return books.Domain{}, result.Error
		}

	}

	return domain, nil
}

func NewMysqlImageBookRepository(conn *gorm.DB) books.ImageBooksRepository {
	return &ImageBooksRepository{
		Conn: conn,
	}
}
