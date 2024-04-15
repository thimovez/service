package comment

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/repo/postgres/comment"
)

type UseCaseComment struct {
	iImageRepo comment.CommentRepository
}

func New(i comment.CommentRepository) *UseCaseComment {
	return &UseCaseComment{
		iImageRepo: i,
	}
}

func (u *UseCaseComment) CreateComment(c entity.Comment) error {
	err := u.iImageRepo.CreateComment(c)
	if err != nil {
		return err
	}

	return nil
}
