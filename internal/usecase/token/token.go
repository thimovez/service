package token

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/repo/postgres/user"
	"github.com/thimovez/service/internal/usecase/token/tokenapi"
	"time"
)

type TokenService interface {
	GenerateAccessToken(a entity.LoginRes) (accessToken string, err error)
	GenerateRefreshToken(a entity.LoginRes) (refreshToken string, err error)
	VerifyAccessToken(tokenString string) error
	VerifyRefreshToken(tokenString string) error
	Refresh(refreshTokenString string) (t entity.RefreshRes, err error)
}

type TokenUseCase struct {
	jwtProvider       tokenapi.JWTProvider
	accessExpiration  time.Time
	refreshExpiration time.Time
	accessSecret      string
	refreshSecret     string
	userRepo          user.UserRepository
}

const minSecretKeySize = 3

func New(
	jwtProvider tokenapi.JWTProvider,
	accessExp time.Time,
	refreshExp time.Time,
	secret string,
	userRepo user.UserRepository,
) *TokenUseCase {
	//TODO add err to New function
	//if secret < minSecretKeySize {
	//	return fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	//}
	return &TokenUseCase{
		jwtProvider:       jwtProvider,
		accessExpiration:  accessExp,
		refreshExpiration: refreshExp,
		accessSecret:      secret,
		refreshSecret:     secret,
		userRepo:          userRepo,
	}
}

func (t *TokenUseCase) GenerateAccessToken(a entity.LoginRes) (accessToken string, err error) {
	claims := map[string]interface{}{
		"userID":   a.User.ID,
		"username": a.User.Username,
		"exp":      t.accessExpiration.Unix(),
	}

	token, err := t.jwtProvider.CreateToken(claims)
	if err != nil {
		return accessToken, err
	}

	accessToken, err = t.jwtProvider.SignToken(token, []byte(t.accessSecret))
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}

func (t *TokenUseCase) GenerateRefreshToken(a entity.LoginRes) (refreshToken string, err error) {
	claims := map[string]interface{}{
		"userID": a.User.ID,
		"exp":    t.refreshExpiration.Unix(),
	}

	token, err := t.jwtProvider.CreateToken(claims)
	if err != nil {
		return "", err
	}

	refreshToken, err = t.jwtProvider.SignToken(token, []byte(t.refreshSecret))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (t *TokenUseCase) VerifyAccessToken(tokenString string) error {
	err := t.jwtProvider.VerifyToken(tokenString, []byte(t.accessSecret))
	if err != nil {
		return err
	}

	return err
}

func (t *TokenUseCase) VerifyRefreshToken(refreshToken string) error {
	err := t.jwtProvider.VerifyToken(refreshToken, []byte(t.refreshSecret))
	if err != nil {
		return err
	}

	return nil
}

func (t *TokenUseCase) Refresh(refreshTokenString string) (tokens entity.RefreshRes, err error) {
	err = t.jwtProvider.VerifyToken(refreshTokenString, []byte(t.refreshSecret))
	if err != nil {
		return tokens, err
	}

	claims, err := t.jwtProvider.ExtractClaims(refreshTokenString, []byte(t.refreshSecret))
	if err != nil {
		return tokens, err
	}

	t.userRepo.
	// TODO Get id from claims and check if user exist with this id
	// TODO If exist create new access token and refresh token with claims received from ExtractClaims

	return
}
