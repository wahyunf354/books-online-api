package users

type UserRegister struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Role         int8   `json:"role"`
	Birth        string `json:"birth"`
	Gender       string `json:"gender"`
	Password     string `json:"password"`
	ConfPassword string `json:"confpassword"`
}
