package entity

import "time"

type Credentials struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID string `json:"id"`
	Credentials
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegistrationReq struct {
	ID string `json:"id"`
	Credentials
	Role string `json:"role"`
}

type UserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
