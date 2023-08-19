package entity

type UserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Image struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ImagePath string `json:"image_path"`
	ImageURL  string `json:"image_url"`
}

type LoginResponse struct {
	AccessToken string
}
