package image_repo

import (
	"database/sql"
	"github.com/thimovez/service/internal/entity"
	"log"
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

func (i *ImageRepo) GetImages() (images []entity.Image, err error) {
	q := `SELECT * FROM images`

	row, err := i.db.Query(q)
	if err != nil {
		return
	}

	for row.Next() {
		var img entity.Image
		err := row.Scan(&img.ID, &img.UserID, &img.ImagePath, &img.ImageURL)
		if err != nil {
			log.Fatal(err)
		}
		images = append(images, img)
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return images, nil
}
