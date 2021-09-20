package books

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title string `gorm:"not null"`
	Books []Book `gorm:"foreignKey:CategoryId;references:ID"`
}

type BookType struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Unit  string `gorm:"not null"` // unit digunakan ketika disewa itu untuk per hari/bulan/selamanya
	Books []Book `gorm:"foreignKey:BookTypeId;references:ID"`
}

type ImageBooks struct {
	gorm.Model
	Url    string `gorm:"not null"`
	BookId int    `gorm:"not null"`
	Book   Book
}

type BookDetails struct {
	gorm.Model
	UrlBook     string `gorm:"not null"`
	PageCount   int
	Description string `gorm:"not null"`
	BookId      string `gorm:"not null;unique"`
}

type Book struct {
	gorm.Model
	Title      string       `gorm:"not null"`
	BookTypeId int          `gorm:"not null"`
	CategoryId int          `gorm:"not null"`
	Price      int          `gorm:"not null"`
	UserId     int          `gorm:"not null"`
	BookType   BookType     `gorm:"foreignKey:BookTypeId"`
	Category   Category     `gorm:"foreignKey:CategoryId"`
	Images     []ImageBooks `gorm:"foreignKey:BookId"`
}
