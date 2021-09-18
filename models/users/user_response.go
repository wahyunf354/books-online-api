package users

type UserRegister struct {
	UserName     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Role         int8   `json:"role"`
	Birth        string `json:"birth`
	Password     string `json:"password"`
	ConfPassword string `json:"confpassword"`
}
