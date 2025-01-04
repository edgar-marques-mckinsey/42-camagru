package models

import (
	"backend/utils"
	"time"
)

type Image struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Data      string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateImage(id int, image []byte) error {
	db := utils.GetDB()
	_, err := db.Exec(`
			INSERT INTO images (data, user_id)
			VALUES ($1, $2)
		`, image, id)

	return err
}

func GetImages() ([]Image, error) {
	db := utils.GetDB()
	rows, err := db.Query(`
			SELECT id, user_id, data, created_at
			FROM images
		`)
	if err != nil {
		return nil, err
	}

	images := []Image{}
	for rows.Next() {
		var image Image
		err := rows.Scan(&image.ID, &image.UserID, &image.Data, &image.CreatedAt)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}

	return images, nil
}
