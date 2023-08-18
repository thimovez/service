package repo

import (
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) SaveUser(id int64, username, password string) error {
	q := `INSERT INTO users (id, username, password_hash) VALUES ($1, $2, $3)`

	_, err := u.db.Exec(q, id, username, password)
	if err != nil {
		return err
	}

	return nil
}
