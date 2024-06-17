package entity

type Image struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	ImagePath string `json:"image_path"`
	ImageURL  string `json:"image_url"`
}
