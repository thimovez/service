package image

import (
	"github.com/google/uuid"
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
	id := uuid.New().String()
	image.ID = id

	err := u.image.SaveImage(image)
	if err != nil {
		return err
	}

	return nil
}

func (u *ImageUseCase) GetImages() (images []entity.Image, err error) {
	images, err = u.image.GetImages()
	if err != nil {
		return nil, err
	}
	return images, nil
}
