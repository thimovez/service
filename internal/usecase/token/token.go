package token

import (
	"github.com/thimovez/service/internal/providers/auth"
)

type UseCaseAuth struct {
	jwtProvider auth.JWTProvider
}

func New(jwtProvider auth.JWTProvider) *UseCaseAuth {
	return &UseCaseAuth{jwtProvider: jwtProvider}
}

func (t *UseCaseAuth) GenerateAccessToken(userID string) (accessToken string, err error) {
	token, err := t.jwtProvider.CreateToken(userID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *UseCaseAuth) VerifyAccessToken(tokenString string) (map[string]interface{}, error) {
	claims, err := t.jwtProvider.VerifyToken(tokenString)
	if err != nil {
		return claims, err
	}

	return claims, err
}
