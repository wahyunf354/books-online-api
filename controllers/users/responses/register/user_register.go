package register

import (
	"books_online_api/business/users"
	"time"
)

type UserRegister struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) UserRegister {
	return UserRegister{
		Id: domain.Id,
		FirstName: domain.FirstName,
		Email: domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
