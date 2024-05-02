package entity

type AuthorizationRes struct {
	User  UserRes `json:"user"`
	Token Token   `json:"token"`
}

type AuthorizationReq struct {
	User UserRegistrationReq `json:"user" validate:"required"`
}
