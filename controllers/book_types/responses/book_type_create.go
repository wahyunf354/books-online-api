package responses

import (
	"books_online_api/business/book_types"
	"gorm.io/gorm"
	"time"
)

type BookTypeResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Unit string `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain book_types.Domain) BookTypeResponse {
	return BookTypeResponse{
		Id: domain.Id,
		Name: domain.Name,
		Unit: domain.Unit,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
