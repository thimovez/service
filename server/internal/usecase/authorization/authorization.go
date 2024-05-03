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
	VerifyLoginData(c context.Context, a entity.AuthorizationReq) (validData entity.AuthorizationReq, err error)
	VerifyRegistrationData(c context.Context, user entity.UserRegistrationReq) (err error)
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

func (u *AuthUseCase) VerifyLoginData(c context.Context, a entity.AuthorizationReq) (validData entity.AuthorizationReq, err error) {
	hashedPassword, err := u.iUserRepo.GetPassword(c, a.User.Username)
	if err != nil {
		return
	}

	err = u.iBcryptProvider.ComparePassword([]byte(hashedPassword), []byte(a.User.Password))
	if err != nil {
		return
	}

	id, err := u.iUserRepo.GetID(c, a.User.Username)
	if err != nil {
		return
	}

	validData.User.ID = id
	validData.User.Username = a.User.Username

	return validData, nil
}

func (u *AuthUseCase) VerifyRegistrationData(c context.Context, user entity.UserRegistrationReq) (err error) {
	err = u.iUserRepo.GetUsername(c, user.Username)
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

	err = u.iUserRepo.SaveUser(c, user)
	if err != nil {
		return
	}

	return nil
}
