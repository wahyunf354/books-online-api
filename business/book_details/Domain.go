package book_details

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	Id          int
	UrlBook     string
	PageCount   int
	Description string
	BookId      int

	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
