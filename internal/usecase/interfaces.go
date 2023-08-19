package usecase

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/thimovez/service/internal/entity"
	"time"
)

type (
	UserService interface {
		Login(user entity.UserRequest) (accessToken string, err error)
	}

	TokenService interface {
		GenerateAccessToken(userID string, expiration time.Time) (string, error)
		VerifyAccessToken(tokenString string) (*jwt.Token, error)
	}

	UserRepo interface {
		SaveUser(user entity.UserRequest) error
		CheckUsername(username string) error
	}

	ImageRepo interface {
		SaveImage(image entity.Image) error
		GetImages() error
	}
)
