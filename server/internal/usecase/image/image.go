package image

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/usecase/authorization/uuidapi"
	"github.com/thimovez/service/internal/usecase/repo/postgres/image"
)

type UseCaseImage struct {
	iImageRepo    image.ImageRepository
	iUUIDProvider uuidapi.UUIDProvider
}

func New(i image.ImageRepository, up uuidapi.UUIDProvider) *UseCaseImage {
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
