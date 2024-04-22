package entity

type AuthorizationRes struct {
	User   UserResponse `json:"user"`
	Tokens Tokens       `json:"tokens"`
}

type AuthorizationReq struct {
	User UserRequest `json:"user"`
}
