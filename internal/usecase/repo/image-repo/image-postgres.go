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

func (i *ImageRepo) SaveImage(image entity.Image) error {
	q := `INSERT INTO images (id, user_id, image_path, image_url) VALUES ($1, $2, $3, $4)`

	_, err := i.db.Exec(q, image.ID, image.UserID, image.ImagePath, image.ImageURL)
	if err != nil {
		return err
	}

	return nil
}

func (i *ImageRepo) GetImages() (images []string, err error) {
	q := `SELECT * FROM images`

	row, err := i.db.Query(q)
	if err != nil {
		return
	}

	images, err = row.Columns()
	if err != nil {
		return
	}

	return images, nil
}
