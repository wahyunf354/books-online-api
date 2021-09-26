package books

import (
	"books_online_api/business/books"
	"books_online_api/drivers/databases/book_types"
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Title      string        `gorm:"not null"`
	BookTypeId int           `gorm:"not null"`
	CategoryId int           `gorm:"not null"`
	Price      int           `gorm:"not null"`
	UserId     int           `gorm:"not null"`
	BookType   *book_types.BookType     `gorm:"foreignKey:BookTypeId"`
	//Category   *Category     `gorm:"foreignKey:CategoryId"`
	//Images     []*ImageBooks `gorm:"foreignKey:BookId"`
}

func FromDomain(domain books.Domain) Books {
	return Books{
		Title: domain.Title,
		BookTypeId: domain.BookTypeId,
		Price: domain.Price,
		UserId: domain.UserId,
	}
}
