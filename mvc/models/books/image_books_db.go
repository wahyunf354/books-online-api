package books

import (
	"books_online_api/business/image_books"
	"gorm.io/gorm"
	"time"
)

type ImageBooks struct {
	Id int `gorm:"primaryKey"`
	Url    string `gorm:"not null"`
	BookId uint   `gorm:"not null"`
	Book   *Book

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func FormDomain(domain image_books.Domain) ImageBooks {
	return ImageBooks{
		Url:       domain.UrlCover,
		BookId:    0,
		Book:      nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}
