package books

import "gorm.io/gorm"

type BookDetail struct {
	gorm.Model
	UrlBook     string `gorm:"not null"`
	PageCount   int
	Description string `gorm:"not null"`
	BookId      uint   `gorm:"not null;unique"`
}
