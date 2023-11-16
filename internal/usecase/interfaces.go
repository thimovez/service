package usecase

import (
	"github.com/thimovez/service/internal/entity"
)

// TODO remove type interface from file inteface. And delete them
type (
	ImageRepo interface {
		SaveImage(image entity.Image) error
		GetImages() (images []entity.Image, err error)
	}
)
