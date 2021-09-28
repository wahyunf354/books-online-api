package books

import (
	"books_online_api/business/books"
	"books_online_api/drivers/databases/book_types"
	"books_online_api/drivers/databases/image_books"
	"gorm.io/gorm"
	"time"
)

type Book struct {
	Id int `gorm:"primaryKey"`
	Title      string        `gorm:"not null"`
	BookTypeId int           `gorm:"not null"`
	CategoryId int
	Price      int           `gorm:"not null"`
	UserId     int           `gorm:"not null"`
	BookType   *book_types.BookType     `gorm:"foreignKey:BookTypeId"`
	//Category   *Category     `gorm:"foreignKey:CategoryId"`
	Images     []*image_books.ImageBooks `gorm:"foreignKey:BookId"`
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}


func FromDomain(domain books.Domain) Book {
	return Book{
		Title: domain.Title,
		BookTypeId: domain.BookTypeId,
		Price: domain.Price,
		UserId: domain.UserId,
	}
}

func (book  Book) ToDomain(domain books.Domain) books.Domain {
	return books.Domain{
		Id: book.Id,
		Price: book.Price,
		Title: book.Title,
		BookTypeId: book.BookTypeId,
		UserId: book.UserId,

		PageCount: domain.PageCount,
		Description: domain.Description,
		UrlBook: domain.UrlBook,
		FileCover: domain.FileCover,
		UrlCover: domain.UrlCover,

		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
		DeletedAt: book.DeletedAt,
	}
}

