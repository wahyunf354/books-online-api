package requests

import "books_online_api/business/users"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (userLogin *UserLogin) ToDomain() users.Domain {
	return users.Domain{
		Email: userLogin.Email,
		Password: userLogin.Password,
	}
}

