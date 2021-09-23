package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `json:"-"`
	Role      int8   `gorm:"not null;default:1"` // 1 for reader, 2 for writer
	Birth     time.Time
	Gender    string
}
