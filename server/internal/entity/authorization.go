package entity

type AuthorizationReq struct {
	User Credentials
}

type LoginRes struct {
	User  UserLoginRes `json:"user"`
	Token Token        `json:"token"`
}

type LoginReq struct {
	User UserLoginReq `json:"user"`
}
