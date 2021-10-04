package payment_methods

import (
	"books_online_api/business/payment_methods"
	"context"
	"gorm.io/gorm"
)

type PaymentMethodsRepository struct {
	Conn *gorm.DB
}


func NewMysqlPaymentMethodsRepository (conn *gorm.DB) payment_methods.Repository {
	return PaymentMethodsRepository{Conn: conn}
}

func (p PaymentMethodsRepository) CreatePaymentMethod(ctx context.Context, domain payment_methods.Domain) (payment_methods.Domain, error) {
	var paymentMethod PaymentMethod
	
	paymentMethod = FromDomain(domain)
	
	resultDb := p.Conn.Create(&paymentMethod)
	
	if resultDb.Error != nil {
		return payment_methods.Domain{}, resultDb.Error
	}
	
	return paymentMethod.ToDomain(), nil
}

func (p PaymentMethodsRepository) GetOnePaymentMethod(ctx context.Context, domain payment_methods.Domain) (payment_methods.Domain, error) {
	var paymentMethod PaymentMethod

	resultPaymentMethod := p.Conn.First(&paymentMethod, domain.Id)

	if resultPaymentMethod.Error != nil {
		return payment_methods.Domain{}, resultPaymentMethod.Error
	}

	return paymentMethod.ToDomain(), nil
}
