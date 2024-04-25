package authorization

import (
	"context"
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/authorization/bcryptapi"
	"github.com/thimovez/service/internal/usecase/authorization/uuidapi"
	"github.com/thimovez/service/internal/usecase/repo/postgres/user"
	"github.com/thimovez/service/internal/usecase/token"
)

type AuthUserService interface {
	VerifyLoginData(c context.Context, a entity.AuthorizationReq) (res entity.AuthorizationRes, err error)
	VerifyRegistrationData(c context.Context, user entity.UserRequest) (err error)
}

// AuthUserUseCase - prefix i means that this is an interface
type AuthUserUseCase struct {
	iUserRepo       user.UserRepository
	iTokenService   token.TokenService
	iBcryptProvider bcryptapi.BcryptProvider
	iUUIDProvider   uuidapi.UUIDProvider
}

func New(
	u user.UserRepository,
	t token.TokenService,
	up uuidapi.UUIDProvider,
	bp bcryptapi.BcryptProvider,
) *AuthUserUseCase {
	return &AuthUserUseCase{
		iUserRepo:       u,
		iTokenService:   t,
		iBcryptProvider: bp,
		iUUIDProvider:   up,
	}
}

func (u *AuthUserUseCase) VerifyLoginData(c context.Context, a entity.AuthorizationReq) (res entity.AuthorizationRes, err error) {
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

	a.User.ID = id

	accessToken, err := u.iTokenService.GenerateAccessToken(a)
	if err != nil {
		return
	}

	refreshToken, err := u.iTokenService.GenerateRefreshToken(a)
	if err != nil {
		return
	}

	res.Tokens.AccessToken = accessToken
	res.Tokens.RefreshToken = refreshToken
	res.User.ID = id
	res.User.Username = a.User.Username

	return res, nil
}

func (u *AuthUserUseCase) VerifyRegistrationData(c context.Context, user entity.UserRequest) (err error) {
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
