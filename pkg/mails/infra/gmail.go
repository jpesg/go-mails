package mails

import (
	sharedD "go-mails/pkg/shared/domain"
	"log"
	"net/smtp"
)

func Send(content string, from string, to string, pass string) (bool, error) {
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		content

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false, sharedD.New("Error sending the email")
	}
	return true, nil
}
