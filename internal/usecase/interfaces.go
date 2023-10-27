package usecase

import (
	"github.com/thimovez/service/internal/entity"
)

type (
	UserService interface {
		Login(user entity.UserRequest) (accessToken string, err error)
	}

	TokenService interface {
		GenerateAccessToken(userID string) (string, error)
		VerifyAccessToken(tokenString string) (map[string]interface{}, error)
	}

	UserRepo interface {
		SaveUser(user entity.UserRequest) error
		CheckUsername(username string) error
	}

	ImageRepo interface {
		SaveImage(image entity.Image) error
		GetImages() (images []entity.Image, err error)
	}
)
