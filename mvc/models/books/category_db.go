package books

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title string  `gorm:"not null"`
	Books []*Book `gorm:"foreignKey:CategoryId;references:ID"`
}
