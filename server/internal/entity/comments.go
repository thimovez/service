package entity

type Comment struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Content  string `json:"content"`
	ParentID string `json:"parent_id"`
}
