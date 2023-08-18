package postgres

import (
	"database/sql"
)

type Storage interface {
	SaveUser(id int64, username, password string) error
	//SaveImage() error
	//GetImage() error
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{db: db}
}

func (s *storage) SaveUser(id int64, username, password string) error {
	q := `INSERT INTO users (id, username, password_hash), VALUES (?, ?, ?)`

	_, err := s.db.Exec(q, id, username, password)
	if err != nil {
		return err
	}

	return nil
}
