package image_repo

import (
	"database/sql"
	"github.com/thimovez/service/internal/entity"
)

type ImageRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *ImageRepo {
	return &ImageRepo{db}
}

//TODO repo
func (u *ImageRepo) SaveImage(image entity.Image) error {

	return nil
}

// TODO
func (u *ImageRepo) GetImages() error {
	return nil
}
