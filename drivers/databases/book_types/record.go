package book_types

import (
	"books_online_api/business/book_types"
	"gorm.io/gorm"
	"time"
)

type BookType struct {
	Id int `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Unit string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (book_type *BookType) ToDomain() book_types.Domain {
	return book_types.Domain{
		Id: book_type.Id,
		Name: book_type.Name,
		Unit: book_type.Unit,
		CreatedAt: book_type.CreatedAt,
		UpdatedAt: book_type.UpdatedAt,
	}
}

func FromDomain(domain book_types.Domain) BookType {
	return BookType{
		Name: domain.Name,
		Unit: domain.Unit,
	}
}
