package requests

import "time"

type UserRegister struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Birth        time.Time `json:"birth"`
	Gender       string    `json:"gender"`
	ConfPassword string    `json:"confpassword"`
}
