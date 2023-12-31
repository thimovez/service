package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const minSecretKeySize = 3

type JWTProvider interface {
	CreateToken(userID string) (string, error)
	VerifyToken(tokenString string) (map[string]interface{}, error)
}

type JWTProviderImpl struct {
	secretKey  string
	expiration time.Time
}

func NewJWTProvider(secretKey string, expiration time.Time) (JWTProvider, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTProviderImpl{
		secretKey:  secretKey,
		expiration: expiration,
	}, nil
}

func (provider *JWTProviderImpl) CreateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    provider.expiration.Unix(),
	})

	tokenString, err := token.SignedString([]byte(provider.secretKey))
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func (provider *JWTProviderImpl) VerifyToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure to validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used for signing
		return []byte(provider.secretKey), nil
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
