package token

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/token/tokenapi"
	"time"
)

type TokenService interface {
	GenerateAccessToken(a entity.AuthorizationReq) (accessToken string, err error)
	GenerateRefreshToken(a entity.AuthorizationReq) (refreshToken string, err error)
	VerifyAccessToken(tokenString string) (map[string]interface{}, error)
}

type TokenUseCase struct {
	jwtProvider       tokenapi.JWTProvider
	accessExpiration  time.Time
	refreshExpiration time.Time
}

func New(jwtProvider tokenapi.JWTProvider, accessExp time.Time, refreshExp time.Time) *TokenUseCase {
	return &TokenUseCase{
		jwtProvider:       jwtProvider,
		accessExpiration:  accessExp,
		refreshExpiration: refreshExp,
	}
}

func (t *TokenUseCase) GenerateAccessToken(a entity.AuthorizationReq) (accessToken string, err error) {
	claims := map[string]interface{}{
		"userID":   a.User.ID,
		"username": a.User.Username,
		"exp":      t.accessExpiration.Unix(),
	}

	token, err := t.jwtProvider.CreateToken(claims)
	if err != nil {
		return accessToken, err
	}

	accessToken, err = t.jwtProvider.SignToken(token)
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}

func (t *TokenUseCase) GenerateRefreshToken(a entity.AuthorizationReq) (refreshToken string, err error) {
	claims := map[string]interface{}{
		"userID": a.User.ID,
		"exp":    t.refreshExpiration.Unix(),
	}

	token, err := t.jwtProvider.CreateToken(claims)
	if err != nil {
		return refreshToken, err
	}

	refreshToken, err = t.jwtProvider.SignToken(token)
	if err != nil {
		return refreshToken, err
	}

	return refreshToken, nil
}

func (t *TokenUseCase) VerifyAccessToken(tokenString string) (map[string]interface{}, error) {
	claims, err := t.jwtProvider.VerifyToken(tokenString)
	if err != nil {
		return claims, err
	}

	return claims, err
}
