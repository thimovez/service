package usecase

import (
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo UserRepo
}

func New(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (u *UserUseCase) Login(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//TODO add id
	var id int64
	err = u.repo.SaveUser(id, username, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}
