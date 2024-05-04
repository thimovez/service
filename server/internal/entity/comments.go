package entity

type Comment struct {
	ID      uint64        `json:"id"`
	Author  User          `json:"author"`
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
	Author  User    `json:"author"`
	Comment Comment `json:"comment"`
}
