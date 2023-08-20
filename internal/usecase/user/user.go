package user

import (
	"github.com/google/uuid"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase"
	"golang.org/x/crypto/bcrypt"
)

type UseCaseUser struct {
	iUserRepo     usecase.UserRepo
	iTokenService usecase.TokenService
}

func New(u usecase.UserRepo, t usecase.TokenService) *UseCaseUser {
	return &UseCaseUser{
		iUserRepo:     u,
		iTokenService: t,
	}
}

func (u *UseCaseUser) Login(user entity.UserRequest) (accessToken string, err error) {
	err = u.iUserRepo.CheckUsername(user.Username)
	if err != nil {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = string(hashedPassword)

	id := uuid.New()
	user.ID = id.String()

	err = u.iUserRepo.SaveUser(user)
	if err != nil {
		return
	}

	accessToken, err = u.iTokenService.GenerateAccessToken(user.ID)
	if err != nil {
		return
	}

	return accessToken, nil
}
