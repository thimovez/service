package user

import (
	"context"
	"database/sql"
	"github.com/thimovez/service/internal/entity"
	"log"
)

type UserRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepo {
	return &UserRepo{
		db,
	}
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
// TODO rename function to get username
func (u *UserRepo) CheckUsername(username string) error {
	q := `SELECT ( username ) FROM users WHERE username = $1`

	var user sql.NullString
	// TODO add context
	err := u.db.QueryRow(q, username).Scan(&user)
	if !user.Valid {
		if err == sql.ErrNoRows {
			return nil
		}
	}

	return err
}

// GetPassword - return hashed password by username.
func (u *UserRepo) GetPassword(username string) (hashedPassword string, error error) {
	qGetPassword := `SELECT (password_hash) FROM users WHERE username = $1`
	err := u.db.QueryRowContext(context.TODO(), qGetPassword, username).Scan(&hashedPassword)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", username)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created on %s\n", username)
	}

	return hashedPassword, nil
}
