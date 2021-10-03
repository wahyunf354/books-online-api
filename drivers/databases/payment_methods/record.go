package payment_methods

import (
	"books_online_api/business/payment_methods"
	"gorm.io/gorm"
	"time"
)

type PaymentMethod struct {
	Id      int    `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Fee     int
	Address string `gorm:"not null"`
	Author  string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (p PaymentMethod) ToDomain() payment_methods.Domain {
	return payment_methods.Domain{
		Id:        p.Id,
		Name:      p.Name,
		Fee:       p.Fee,
		Address:   p.Address,
		Author:    p.Author,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}

func FromDomain(domain payment_methods.Domain) PaymentMethod {
	return PaymentMethod{
		Name:      domain.Name,
		Fee:       domain.Fee,
		Address:   domain.Address,
		Author:    domain.Author,
	}
}
