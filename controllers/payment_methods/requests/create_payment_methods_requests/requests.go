package create_payment_methods_requests

import (
	"books_online_api/business/payment_methods"
)

type CreatePaymentMethodRequest struct {
	Name    string `json:"name"`
	Fee     int `json:"fee"`
	Address string `json:"address"`
	Author  string `json:"author"`

}

func (c CreatePaymentMethodRequest) ToDomain() payment_methods.Domain {
	return payment_methods.Domain{
		Name:      c.Name,
		Fee:       c.Fee,
		Address:   c.Address,
		Author:    c.Author,
	}
}

func FromDomain(domain payment_methods.Domain) CreatePaymentMethodRequest {
	return CreatePaymentMethodRequest{
		Name:    domain.Name,
		Fee:     domain.Fee,
		Address: domain.Address,
		Author:  domain.Author,
	}
}
