package user

import (
	"github.com/google/uuid"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func (u *UserUseCase) Login(user entity.UserRequest) (accessToken string, err error) {
	err = u.repo.CheckUsername(user.Username)
	if err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	user.Password = string(hashedPassword)

	id := uuid.New()
	user.ID = id.String()

	err = u.repo.SaveUser(user)
	if err != nil {
		return "", nil
	}

	expiration := time.Now().Add(time.Hour * 12)
	accessToken, err = u.token.GenerateAccessToken(user.ID, expiration)
	if err != nil {
		return "", nil
	}

	return accessToken, nil
}
