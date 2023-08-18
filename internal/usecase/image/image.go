package image

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase"
)

type ImageUseCase struct {
	image usecase.ImageRepo
}

func New(i usecase.ImageRepo) *ImageUseCase {
	return &ImageUseCase{
		image: i,
	}
}

func (u *ImageUseCase) SaveImage(image entity.Image) error {
	return nil
}

func (u *ImageUseCase) GetImages() error {
	return nil
}
