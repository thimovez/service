package user

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/providers/bcrypt"
	"github.com/thimovez/service/internal/providers/uuid"
	"github.com/thimovez/service/internal/usecase"
)

// UseCaseUser - prefix i means that this is an interface
type UseCaseUser struct {
	iUserRepo       usecase.UserRepo
	iTokenService   usecase.TokenService
	iBcryptProvider bcrypt.BcryptProvider
	iUUIDProvider   uuid.UUIDProvider
}

// TODO сделать передачу лишних агрументов как опции
func New(u usecase.UserRepo, t usecase.TokenService, up uuid.UUIDProvider, bp bcrypt.BcryptProvider) *UseCaseUser {
	return &UseCaseUser{
		iUserRepo:       u,
		iTokenService:   t,
		iBcryptProvider: bp,
		iUUIDProvider:   up,
	}
}

func (u *UseCaseUser) Login(user entity.UserRequest) (accessToken string, err error) {
	hashedPassword, err := u.iUserRepo.GetPassword(user.Username)
	if err != nil {
		return
	}

	err = u.iBcryptProvider.ComparePassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return
	}

	accessToken, err = u.iTokenService.GenerateAccessToken(user.ID)
	if err != nil {
		return
	}

	return accessToken, nil
}

func (u *UseCaseUser) Registration(user entity.UserRequest) (err error) {
	err = u.iUserRepo.CheckUsername(user.Username)
	if err != nil {
		return
	}

	hashedPassword, err := u.iBcryptProvider.HashPassword(user.Password)
	if err != nil {
		return
	}

	user.Password = string(hashedPassword)

	id := u.iUUIDProvider.CreateStringUUID()
	user.ID = id

	err = u.iUserRepo.SaveUser(user)
	if err != nil {
		return
	}

	return nil
}
