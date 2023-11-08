package image

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/providers/uuid"
	"github.com/thimovez/service/internal/usecase"
)

type UseCaseImage struct {
	iImageRepo    usecase.ImageRepo
	iUUIDProvider uuid.UUIDProvider
}

func New(i usecase.ImageRepo, up uuid.UUIDProvider) *UseCaseImage {
	return &UseCaseImage{
		iImageRepo:    i,
		iUUIDProvider: up,
	}
}

func (u *UseCaseImage) SaveImage(image entity.Image) error {
	id := u.iUUIDProvider.CreateStringUUID()
	image.ID = id

	err := u.iImageRepo.SaveImage(image)
	if err != nil {
		return err
	}

	return nil
}

func (u *UseCaseImage) GetImages() (images []entity.Image, err error) {
	images, err = u.iImageRepo.GetImages()
	if err != nil {
		return
	}
	return images, nil
}
