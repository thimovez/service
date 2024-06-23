package tokenapi

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider interface {
	CreateToken(claims map[string]interface{}) (token *jwt.Token, err error)
	SignToken(token *jwt.Token, secret []byte) (signedToken string, err error)
	VerifyToken(tokenString string, secret []byte) error
	ExtractClaims(tokenStr string, secret []byte) (claims jwt.MapClaims, err error)
}

type JWTProviderImpl struct{}

func New() (JWTProvider, error) {
	return &JWTProviderImpl{}, nil
}

func (j *JWTProviderImpl) CreateToken(claims map[string]interface{}) (token *jwt.Token, err error) {
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))

	return token, nil
}

func (j *JWTProviderImpl) SignToken(token *jwt.Token, secret []byte) (signedToken string, err error) {
	signedToken, err = token.SignedString(secret)
	if err != nil {
		return "", nil
	}

	return signedToken, nil
}

func (j *JWTProviderImpl) VerifyToken(tokenString string, secret []byte) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure to validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func (j *JWTProviderImpl) ExtractClaims(tokenStr string, secret []byte) (claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return claims, fmt.Errorf("invalid jwt token")
	}

	return claims, nil
}
