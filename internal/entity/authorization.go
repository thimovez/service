package entity

type LoginRes struct {
	User  UserLoginRes `json:"user"`
	Token Token        `json:"token"`
}

type LoginReq struct {
	User UserLoginReq `json:"user"`
}

type RegistrationReq struct {
	User Credentials `json:"user"`
}

type RefreshRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
