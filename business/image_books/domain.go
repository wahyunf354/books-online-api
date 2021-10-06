package image_books

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	Id     int    `gorm:"primaryKey"`
	Url    string `gorm:"not null"`
	BookId int    `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

