package login

import (
	"books_online_api/business/users"
	"time"
)

type UserRegister struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	LastName  string    `json:"last_name"`
	Role      int8      `json:"role"` // 1 for reader, 2 for writer
	Birth     time.Time `json:"birth"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) UserRegister {
	return UserRegister{
		Id:        domain.Id,
		FirstName: domain.FirstName,
		Email:     domain.Email,
		Token:     domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		LastName:  domain.LastName,
		Role:      domain.Role,
		Birth:     domain.Birth,
		Gender:    domain.Gender,
	}
}
