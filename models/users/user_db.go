package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string `json:"-"`
	Role      int8
	Birth     time.Time
	Gender    string
}
