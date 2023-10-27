package image

import (
	"github.com/thimovez/service/internal/entity"
	"github.com/thimovez/service/internal/providers/helpers"
	"github.com/thimovez/service/internal/usecase"
)

type UseCaseImage struct {
	iImageRepo      usecase.ImageRepo
	iHelperProvider helpers.HelperProvider
}

func New(i usecase.ImageRepo, hp helpers.HelperProvider) *UseCaseImage {
	return &UseCaseImage{
		iImageRepo:      i,
		iHelperProvider: hp,
	}
}

func (u *UseCaseImage) SaveImage(image entity.Image) error {
	id := u.iHelperProvider.CreateStringUUID()
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
