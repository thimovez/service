package user

import (
	"database/sql"
	"github.com/thimovez/service/internal/entity"
)

type UserRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) SaveUser(user entity.UserRequest) error {
	q := `INSERT INTO users (id, username, password_hash) VALUES ($1, $2, $3)`

	_, err := u.db.Exec(q, user.ID, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) CheckUsername(username string) error {
	q := `SELECT ( username ) FROM users WHERE username = $1`

	var user sql.NullString
	err := u.db.QueryRow(q, username).Scan(&user)
	if !user.Valid {
		if err == sql.ErrNoRows {
			return nil
		}
	}

	return err
}
