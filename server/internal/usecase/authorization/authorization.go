package authorization

import (
	"context"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/authorization/bcryptapi"
	"github.com/thimovez/service/internal/usecase/authorization/uuidapi"
	"github.com/thimovez/service/internal/usecase/repo/postgres/user"
	"github.com/thimovez/service/internal/usecase/token"
)

type AuthService interface {
	VerifyLoginData(c context.Context, a entity.LoginReq) (validData entity.LoginRes, err error)
	VerifyRegistrationData(c context.Context, user entity.RegistrationReq) (err error)
}

// AuthUseCase - prefix i means that this is an interface
type AuthUseCase struct {
	iUserRepo       user.UserRepository
	iTokenService   token.TokenService
	iBcryptProvider bcryptapi.BcryptProvider
	iUUIDProvider   uuidapi.UUIDProvider
}

func New(
	u user.UserRepository,
	up uuidapi.UUIDProvider,
	bp bcryptapi.BcryptProvider,
) *AuthUseCase {
	return &AuthUseCase{
		iUserRepo:       u,
		iBcryptProvider: bp,
		iUUIDProvider:   up,
	}
}

func (a *AuthUseCase) VerifyLoginData(c context.Context, l entity.LoginReq) (validData entity.LoginRes, err error) {
	hashedPassword, err := a.iUserRepo.GetPassword(c, l.User.Username)
	if err != nil {
		return
	}

	err = a.iBcryptProvider.ComparePassword([]byte(hashedPassword), []byte(l.User.Password))
	if err != nil {
		return
	}

	id, err := a.iUserRepo.GetID(c, l.User.Username)
	if err != nil {
		return
	}

	validData.User.ID = id
	validData.User.Username = l.User.Username

	return validData, nil
}

func (a *AuthUseCase) VerifyRegistrationData(c context.Context, r entity.RegistrationReq) (err error) {
	err = a.iUserRepo.GetUsername(c, r.User.Username)
	if err != nil {
		return
	}

	hashedPassword, err := a.iBcryptProvider.HashPassword(r.User.Password)
	if err != nil {
		return
	}

	id := a.iUUIDProvider.CreateStringUUID()

	u := entity.User{
		ID: id,
		Credentials: entity.Credentials{
			Email:    r.User.Email,
			Username: r.User.Username,
			Password: string(hashedPassword),
		},
	}

	err = a.iUserRepo.SaveUser(c, u)
	if err != nil {
		return
	}

	return nil
}
