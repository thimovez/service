package comment

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/repo/postgres/comment"
)

type Comment interface {
	CreateComment(c entity.Comment) error
}

type UseCaseComment struct {
	iImageRepo comment.CommentRepository
}

func New(i comment.CommentRepository) *UseCaseComment {
	return &UseCaseComment{
		iImageRepo: i,
	}
}

func (u *UseCaseComment) CreateComment(c entity.Comment) error {
	// This checks a comment is a main comment or a parent comment.
	// If parent id is 0 - this mean is main
	//parendID := getParentIDByID(c.ID)

	//if c.ParentID != 0 {
	//	c.ParentID++
	//}

	//if c.ID == "" {
	//	c.ParentID = 0
	//}

	err := u.iImageRepo.Create(c)
	if err != nil {
		return err
	}

	return nil
}
