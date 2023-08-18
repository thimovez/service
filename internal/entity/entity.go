package entity

type UserRequest struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Image struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	ImagePath string `json:"image_path"`
	ImageURL  string `json:"image_url"`
}
