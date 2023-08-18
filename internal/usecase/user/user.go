package user

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo  usecase.UserRepo
	token usecase.TokenService
}

func New(r usecase.UserRepo, t usecase.TokenService) *UserUseCase {
	return &UserUseCase{
		repo:  r,
		token: t,
	}
}

func (u *UserUseCase) Login(user entity.UserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	err = u.repo.SaveUser(user)
	if err != nil {
		return err
	}

	return nil
}
