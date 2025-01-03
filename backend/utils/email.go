package utils

import "log"

func SendEmail(toEmail, subject, body string) error {
	log.Println("toEmail:", toEmail)
	log.Println("subject:", subject)
	log.Println("body:", body)

	return nil
}
