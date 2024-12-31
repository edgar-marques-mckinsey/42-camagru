package validity

import (
	"errors"

	"backend/utils"
)

func DoesUsernameExists(username string) bool {
	db := utils.GetDB()
	row := db.QueryRow(`
			SELECT COUNT(*)
			FROM users
			WHERE username = $1
		`, username)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}

func ValidateUsername(username string) error {
	if (len(username) < 3) || (len(username) > 20) {
		return errors.New("username must be between 3 and 20 characters")
	}
	if DoesUsernameExists(username) {
		return errors.New("username already exists")
	}
	return nil
}

func ValidateUser(username, email, password string) error {
	err := ValidateUsername(username)
	if err != nil {
		return err
	}
	return nil
}
