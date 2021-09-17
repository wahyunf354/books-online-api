package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     *string `grom:"unique"`
	Password  string
	Role      int8
	Birth     *time.Time
	Gender    string
}
