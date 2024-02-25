package user

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/thimovez/service/internal/entity"
)

type UserRepository interface {
	SaveUser(c context.Context, eu entity.UserRequest) error
	GetUsername(c context.Context, username string) error
	GetPassword(username string) (hashedPassword string, err error)
	GetUserDataByID(id string) (user *entity.User, err error)
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

	ctx, cancel := context.WithTimeout(c, 30*time.Second)
	defer cancel()

	_, err := u.db.ExecContext(ctx, q, user.ID, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// GetUsername - checks the presence of a username row in the database.
// If username row not present in database function return nil.
func (u *UserRepo) GetUsername(c context.Context, username string) error {
	q := `SELECT ( username ) FROM users WHERE username = $1`

	ctx, cancel := context.WithTimeout(c, 30*time.Second)
	defer cancel()

	var user sql.NullString
	err := u.db.QueryRowContext(ctx, q, username).Scan(&user)
	if !user.Valid {
		if err == sql.ErrNoRows {
			return nil
		}
	}

	return err
}

// GetPassword - return hashed password by username.
func (u *UserRepo) GetPassword(username string) (hashedPassword string, err error) {
	qGetPassword := `SELECT (password_hash) FROM users WHERE username = $1`
	err = u.db.QueryRowContext(context.TODO(), qGetPassword, username).Scan(&hashedPassword)
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

func (u *UserRepo) GetUserDataByID(id string) (user *entity.User, err error) {
	qGetUser := `SELECT (id) FROM users WHERE id = $1`
	var userData entity.User

	rows, err := u.db.Query(qGetUser, id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(userData.ID, userData.Username)
		if err != nil {
			return user, err
		}
	}

	return &userData, nil
}
