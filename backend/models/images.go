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
