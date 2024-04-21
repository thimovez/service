package user

import (
	"context"
	"database/sql"
	"github.com/thimovez/service/internal/entity"
	"log"
)

type UserRepository interface {
	SaveUser(c context.Context, user entity.UserRequest) error
	GetUsername(c context.Context, username string) error
	GetPassword(c context.Context, username string) (hashedPassword string, err error)
	GetID(c context.Context, username string) (id string, err error)
}

type UserRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepo {
	return &UserRepo{
		db,
	}
}

// SaveUser - save user in database and return nil if success.
func (u *UserRepo) SaveUser(c context.Context, user entity.UserRequest) error {
	q := `INSERT INTO users (id, username, password_hash) VALUES ($1, $2, $3)`

	_, err := u.db.ExecContext(c, q, user.ID, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// GetUsername - checks the presence of a username row in the database.
// If username row not present in database function return nil.
func (u *UserRepo) GetUsername(c context.Context, username string) error {
	q := `SELECT ( username ) FROM users WHERE username = $1`

	var user sql.NullString
	err := u.db.QueryRowContext(c, q, username).Scan(&user)
	if !user.Valid {
		if err == sql.ErrNoRows {
			return nil
		}
	}

	return err
}

// GetPassword - return hashed password by username.
func (u *UserRepo) GetPassword(c context.Context, username string) (hashedPassword string, error error) {
	qGetPassword := `SELECT (password_hash) FROM users WHERE username = $1`
	err := u.db.QueryRowContext(c, qGetPassword, username).Scan(&hashedPassword)
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

func (u *UserRepo) GetID(c context.Context, username string) (id string, error error) {
	qGetID := `SELECT (id) FROM users WHERE username = $1`
	err := u.db.QueryRowContext(c, qGetID, username).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", username)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q, account created on %s\n", username)
	}

	return id, nil
}
