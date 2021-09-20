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

type Book struct {
	gorm.Model
	Title      string   `gorm:"not null"`
	BookTypeId int      `gorm:"not null"`
	CategoryId int      `gorm:"not null"`
	Price      int      `gorm:"not null"`
	UserId     int      `gorm:"not null"`
	BookType   BookType `gorm:"foreignKey:BookTypeId"`
	Category   Category `gorm:"foreignKey:CategoryId"`
}
