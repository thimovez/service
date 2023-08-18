package usecase

import "github.com/thimovez/service/internal/entity"

type (
	UserService interface {
		Login(user entity.UserRequest) error
	}

	TokenService interface {
		GenerateAccessToken()
		VerifyAccessToken() error
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
