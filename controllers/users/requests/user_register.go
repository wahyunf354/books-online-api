package requests

import (
	"books_online_api/business/users"
	"time"
)

type UserRegister struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Birth        time.Time `json:"birth"`
	Gender       string    `json:"gender"`
	ConfPassword string    `json:"confpassword"`
	Role         int8      `json:"role"`
}

func (userRegister *UserRegister) ToDomain() users.Domain {
	return users.Domain{
		FirstName:    userRegister.FirstName,
		LastName:     userRegister.LastName,
		Email:        userRegister.Email,
		Password:     userRegister.Password,
		Birth:        userRegister.Birth,
		Gender:       userRegister.Gender,
		ConfPassword: userRegister.ConfPassword,
		Role:         userRegister.Role,
	}
}
