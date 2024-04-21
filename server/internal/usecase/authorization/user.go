package authorization

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/providers/bcrypt"
	"github.com/thimovez/service/internal/providers/uuid"
	"github.com/thimovez/service/internal/usecase/repo/postgres/user"
	"github.com/thimovez/service/internal/usecase/token"
)

type AuthUserService interface {
	VerifyLoginData(user entity.UserRequest) (res *entity.LoginResponse, err error)
	VerifyRegistrationData(user entity.UserRequest) (err error)
}

// AuthUserUseCase - prefix i means that this is an interface
type AuthUserUseCase struct {
	iUserRepo       user.UserRepository
	iTokenService   token.TokenService
	iBcryptProvider bcrypt.BcryptProvider
	iUUIDProvider   uuid.UUIDProvider
}

func New(
	u user.UserRepository,
	t token.TokenService,
	up uuid.UUIDProvider,
	bp bcrypt.BcryptProvider,
) *AuthUserUseCase {
	return &AuthUserUseCase{
		iUserRepo:       u,
		iTokenService:   t,
		iBcryptProvider: bp,
		iUUIDProvider:   up,
	}
}

func (u *AuthUserUseCase) VerifyLoginData(user entity.UserRequest) (res *entity.LoginResponse, err error) {
	hashedPassword, err := u.iUserRepo.GetPassword(user.Username)
	if err != nil {
		return
	}

	err = u.iBcryptProvider.ComparePassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return
	}

	accessToken, err := u.iTokenService.GenerateAccessToken(user.ID)
	if err != nil {
		return
	}

	res.Tokens.AccessToken = accessToken

	return res, nil
}

func (u *AuthUserUseCase) VerifyRegistrationData(user entity.UserRequest) (err error) {
	err = u.iUserRepo.GetUsername(user.Username)
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
