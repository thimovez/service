package entity

type UserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserID struct {
	ID string `json:"id"`
}

type LoginResponse struct {
	Tokens struct {
		AccessToken string `json:"access_token"`
	} `json:"tokens"`
}
