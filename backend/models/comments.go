package models

import (
	"backend/utils"
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ImageID   int       `json:"image_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func GetCommentsCount(imageId int) (int, error) {
	db := utils.GetDB()
	row := db.QueryRow(`
			SELECT COUNT(*)
			FROM comments
			WHERE image_id = $1
		`, imageId)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func CreateComment(userID, imageID int, content string) error {
	db := utils.GetDB()
	_, err := db.Exec(`
			INSERT INTO comments (user_id, image_id, content)
			VALUES ($1, $2, $3)
		`, userID, imageID, content)

	return err
}
