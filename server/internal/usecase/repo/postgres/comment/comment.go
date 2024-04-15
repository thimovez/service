package comment

import (
	"database/sql"
	"github.com/thimovez/service/internal/entity"
)

type CommentRepository interface {
	Create(c entity.Comment) error
	//GetCommentID(c entity.Comment) uint64
}

type CommentRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *CommentRepo {
	return &CommentRepo{db}
}

func (cr *CommentRepo) Create(c entity.Comment) error {
	q := `INSERT INTO comments ( id, user_id, content, parent_id )
		  VALUES ($1, $2, $3, $4)`
	_, err := cr.db.Exec(q, c.ID, c.UserID, c.Content, c.ParentID)
	if err != nil {
		return err
	}

	return nil
}

//func (u *CommentRepo) GetCommentID() (parentID uint64, err error) {
//	// Если возвращается ноль то добавляем единицу
//	q := `SELECT parentID FROM comments`
//
//	rows, err := u.db.Query(q)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer rows.Close()
//
//	var id uint64
//	for rows.Next() {
//		err := rows.Scan(&id)
//		if err != nil {
//			return id, err
//		}
//	}
//
//	if id == 0 {
//		id++
//	}
//	//_, err := u.db.Exec(q, user.ID, user.Username, user.Password)
//	//if err != nil {
//	//	return err
//	//}
//
//	return 0, nil
//}
