package books

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title      string        `gorm:"not null"`
	BookTypeId int           `gorm:"not null"`
	CategoryId int           `gorm:"not null"`
	Price      int           `gorm:"not null"`
	UserId     int           `gorm:"not null"`
	BookType   *BookType     `gorm:"foreignKey:BookTypeId"`
	Category   *Category     `gorm:"foreignKey:CategoryId"`
	Images     []*ImageBooks `gorm:"foreignKey:BookId"`
}
