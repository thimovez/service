package entity

type Comment struct {
	ID      uint64        `json:"id"`
	User    UserID        `json:"user"`
	Content string        `json:"content"`
	Parent  ParentComment `json:"parent"`
}

type ParentComment struct {
	ID string `json:"id"`
}

type CreateCommentReq struct {
	ID      string        `json:"id"`
	Content string        `json:"content"`
	Parent  ParentComment `json:"parent"`
}

type CreateCommentRes struct {
	User    UserID  `json:"user"`
	Comment Comment `json:"comment"`
}
