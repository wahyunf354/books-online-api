package book_details

import (
	"books_online_api/business/books"
	"gorm.io/gorm"
	"time"
)

type BookDetails struct {
	Id int `gorm:"primaryKey"`
	UrlBook string `gorm:"not null"`
	PageCount int
	Description string
	BookId int `gorm:"not null;unique"`

	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func FromDomain(domain books.Domain) BookDetails {
	return BookDetails{
		UrlBook: domain.UrlBook,
		PageCount: domain.PageCount,
		Description: domain.Description,
		BookId: domain.Id,
	}
}

func (bd *BookDetails) ToDomain() books.Domain {
	return books.Domain{
		UrlBook: bd.UrlBook,
		PageCount: bd.PageCount,
		Description: bd.Description,
		BookTypeId: bd.Id,
	}
}

