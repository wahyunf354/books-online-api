package transactions

import (
	"gorm.io/gorm"
	"time"
)

type Transactions struct {
	Id int `gorm:"primaryKey"`
	OrderId int
	PaymentMethodId int
	Status string `gorm:"default=Pending"`
	TotalPrice int
	UserId int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}