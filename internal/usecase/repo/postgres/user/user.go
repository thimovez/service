package user

import (
	"context"
	"database/sql"
	"github.com/thimovez/service/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

// SaveUser - save user in database and return nil if success.
func (u *UserRepo) SaveUser(user entity.UserRequest) error {
	q := `INSERT INTO users (id, username, password_hash) VALUES ($1, $2, $3)`

	_, err := u.db.Exec(q, user.ID, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// CheckUsername - checks the presence of a user in the database.
// If user not present in database function return nil.
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

// ComparePassword - compares the hash from the user and in the database.
// If the hash matches then return nil.
func (u *UserRepo) ComparePassword(username, password string) error {
	qGetPassword := `SELECT (password_hash) FROM users WHERE username = $1`
	var hashedPassword string
	err := u.db.QueryRowContext(context.TODO(), qGetPassword, username).Scan(&hashedPassword)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", username)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created on %s\n", password)
	}

	// remove this part of code to provider
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
