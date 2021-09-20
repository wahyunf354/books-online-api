package books

import "gorm.io/gorm"

type BookDetail struct {
	gorm.Model
	UrlBook     string `gorm:"not null"`
	PageCount   int
	Description string `gorm:"not null"`
	BookId      string `gorm:"not null;unique"`
}
