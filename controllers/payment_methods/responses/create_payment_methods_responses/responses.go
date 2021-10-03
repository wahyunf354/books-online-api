package create_payment_methods_responses

import (
	"books_online_api/business/payment_methods"
	"gorm.io/gorm"
	"time"
)

type PaymentResponses struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Fee int `json:"fee"`
	Address string `json:"address"`
	Author string `json:"author"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain payment_methods.Domain) PaymentResponses {
	return PaymentResponses{
		Id:        domain.Id,
		Name:      domain.Name,
		Fee:       domain.Fee,
		Address:   domain.Address,
		Author:    domain.Author,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
