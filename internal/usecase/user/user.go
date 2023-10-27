package user

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/providers/helpers"
	"github.com/thimovez/service/internal/usecase"
)

type UseCaseUser struct {
	iUserRepo       usecase.UserRepo
	iTokenService   usecase.TokenService
	iHelperProvider helpers.HelperProvider
}

func New(u usecase.UserRepo, t usecase.TokenService, hp helpers.HelperProvider) *UseCaseUser {
	return &UseCaseUser{
		iUserRepo:       u,
		iTokenService:   t,
		iHelperProvider: hp,
	}
}

func (u *UseCaseUser) Login(user entity.UserRequest) (accessToken string, err error) {
	err = u.iUserRepo.CheckUsername(user.Username)
	if err != nil {
		return
	}

	hashedPassword, err := u.iHelperProvider.HashPassword(user.Password)
	if err != nil {
		return
	}

	user.Password = string(hashedPassword)

	id := u.iHelperProvider.CreateStringUUID()
	user.ID = id

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
