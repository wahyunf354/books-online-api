package image_books

import (
	"books_online_api/business/books"
	"gorm.io/gorm"
	"time"
)

type ImageBooks struct {
	Id     int    `gorm:"primaryKey"`
	Url    string `gorm:"not null"`
	BookId int    `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func FromDomain(domain books.Domain) []ImageBooks {
	var imagesBooks []ImageBooks

	for _, v := range domain.UrlCover {
		imagesBooks = append(imagesBooks, ImageBooks{
			Url:    v,
			BookId: domain.Id,
		})
	}

	return imagesBooks
}


