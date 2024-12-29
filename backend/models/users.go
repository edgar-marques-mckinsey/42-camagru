package models

import (
	"time"

	"backend/utils"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func GetUsers() ([]User, error) {
	db := utils.GetDB()
	rows, err := db.Query(`
			SELECT id, username, email, password, created_at
			FROM users
		`)
	if err != nil {
		return nil, err
	}

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUser(id int) (User, error) {
	db := utils.GetDB()
	row := db.QueryRow(`
			SELECT id, username, email, password, created_at
			FROM users
			WHERE id = $1
		`, id)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
