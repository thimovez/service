package entity

type AuthorizationRes struct {
	User   UserResponse `json:"user"`
	Tokens struct {
		AccessToken string `json:"access_token"`
	} `json:"tokens"`
}

type AuthorizationReq struct {
	User UserRequest `json:"user"`
}
