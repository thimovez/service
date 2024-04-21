package token

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/token/tokenapi"
)

type TokenService interface {
	GenerateAccessToken(a entity.AuthorizationReq) (string, error)
	VerifyAccessToken(tokenString string) (map[string]interface{}, error)
}

type TokenUseCase struct {
	jwtProvider tokenapi.JWTProvider
}

func New(jwtProvider tokenapi.JWTProvider) *TokenUseCase {
	return &TokenUseCase{jwtProvider: jwtProvider}
}

func (t *TokenUseCase) GenerateAccessToken(a entity.AuthorizationReq) (accessToken string, err error) {
	token, err := t.jwtProvider.CreateToken(a)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *TokenUseCase) VerifyAccessToken(tokenString string) (map[string]interface{}, error) {
	claims, err := t.jwtProvider.VerifyToken(tokenString)
	if err != nil {
		return claims, err
	}

	return claims, err
}
