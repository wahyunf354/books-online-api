package responses

import (
	"books_online_api/business/users"
	"time"
)

type UserResponse struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{}
}
