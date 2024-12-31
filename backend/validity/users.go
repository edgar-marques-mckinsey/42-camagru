package validity

import (
	"errors"
	"regexp"

	"backend/utils"
)

const EMAIL_PATTERN = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func DoesUsernameExist(username string) bool {
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

func DoesEmailExist(email string) bool {
	db := utils.GetDB()
	row := db.QueryRow(`
			SELECT COUNT(*)
			FROM users
			WHERE email = $1
		`, email)

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
	if DoesUsernameExist(username) {
		return errors.New("username already exists")
	}
	return nil
}

func ValidateEmail(email string) error {
	re := regexp.MustCompile(EMAIL_PATTERN)
	if !re.MatchString(email) {
		return errors.New("email is not valid")
	}
	if DoesEmailExist(email) {
		return errors.New("email already exists")
	}
	return nil
}

func ValidateUser(username, email, password string) error {
	err := ValidateUsername(username)
	if err != nil {
		return err
	}
	err = ValidateEmail(email)
	if err != nil {
		return err
	}
	return nil
}
