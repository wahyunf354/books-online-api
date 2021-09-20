package books

import "gorm.io/gorm"

type BookType struct {
	gorm.Model
	Name  string  `gorm:"not null"`
	Unit  string  `gorm:"not null"` // unit digunakan ketika disewa itu untuk per hari/bulan/selamanya
	Books []*Book `gorm:"foreignKey:BookTypeId;references:ID"`
}
