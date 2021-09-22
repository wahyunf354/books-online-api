package users

import "time"

type UserResponse struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	Role      int8      `json:"role"` // 1 for reader, 2 for writer
	Birth     time.Time `json:"birth"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
