package books

import "gorm.io/gorm"

type ImageBooks struct {
	gorm.Model
	Url    string `gorm:"not null"`
	BookId uint   `gorm:"not null"`
	Book   *Book
}
