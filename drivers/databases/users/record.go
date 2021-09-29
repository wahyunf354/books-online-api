package users

import (
	"books_online_api/business/users"
	"books_online_api/drivers/databases/orders"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id           int    `gorm:"primaryKey"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	HashPassword string
	Role         int8 `gorm:"not null;default:1"` // 1 for reader, 2 for writer
	Birth        time.Time
	Gender       string
	Orders       []*orders.Orders `gorm:"foreignKey:UserId;references:Id"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.FirstName,
		Email:        user.Email,
		HashPassword: user.HashPassword,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Role:         user.Role,
		Birth:        user.Birth,
		Gender:       user.Gender,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:           domain.Id,
		FirstName:    domain.FirstName,
		LastName:     domain.FirstName,
		Email:        domain.Email,
		HashPassword: domain.HashPassword,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		Role:         domain.Role,
		Birth:        domain.Birth,
		Gender:       domain.Gender,
	}
}
