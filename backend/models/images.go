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

func GetUserImages(id int) ([]Image, error) {
	db := utils.GetDB()
	rows, err := db.Query(`
			SELECT id, user_id, data, created_at
			FROM images
			WHERE user_id = $1
		`, id)
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

func GetImage(id int) ([]byte, error) {
	db := utils.GetDB()
	row := db.QueryRow(`
			SELECT data
			FROM images
			WHERE id = $1
		`, id)

	var imageData []byte
	err := row.Scan(&imageData)
	if err != nil {
		return nil, err
	}

	return imageData, nil
}
