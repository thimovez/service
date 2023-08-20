package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UseCaseToken struct {
	secretKey  string
	expiration time.Time
}

func New(s string, e time.Time) *UseCaseToken {
	return &UseCaseToken{s, e}
}

func (t *UseCaseToken) GenerateAccessToken(userID string) (accessToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    t.expiration.Unix(),
	})

	tokenString, err := token.SignedString([]byte(t.secretKey))
	if err != nil {
		return
	}

	return tokenString, nil
}

func (t *UseCaseToken) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure to validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used for signing
		return []byte(t.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
