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

func (bd *BookDetails) ToDomain(domain books.Domain) books.Domain {
	return books.Domain{
		BookDetailId: bd.Id,
		UrlBook: bd.UrlBook,
		PageCount: bd.PageCount,
		Description: bd.Description,

		BookTypeId: domain.BookTypeId,
		Price: domain.Price,
		Id: domain.Id,
		Title: domain.Title,
		UserId: domain.UserId,
		FileCover: domain.FileCover,
		UrlCover: domain.UrlCover,

		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

