package usecase

import (
	"github.com/thimovez/service/internal/entity"
)

// TODO remove type interface from file inteface. And delete them
type (
	TokenService interface {
		GenerateAccessToken(userID string) (string, error)
		VerifyAccessToken(tokenString string) (map[string]interface{}, error)
	}

	UserRepo interface {
		SaveUser(user entity.UserRequest) error
		CheckUsername(username string) error
		GetPassword(username string) (hashedPassword string, err error)
	}

	ImageRepo interface {
		SaveImage(image entity.Image) error
		GetImages() (images []entity.Image, err error)
	}
)
