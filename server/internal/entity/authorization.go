package entity

type AuthorizationRes struct {
	User   UserResponse `json:"user"`
	Tokens Token        `json:"token"`
}

type AuthorizationReq struct {
	User UserRequest `json:"user"`
}
