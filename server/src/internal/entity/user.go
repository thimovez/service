package entity

import "time"

type Credentials struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=6,max=64"`
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

type UserLoginRes struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"required"`
}

type UserLoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
