package mails

import (
	"log"
	"net/smtp"
	"sync"
)

//"gopkg.in/gomail.v2"

func MailSend(content string, from string, to []string, pass string, wg *sync.WaitGroup) {
	log.Printf("sending from  %s to %s", from, to)
	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")
	msg := []byte("Subject: Go!\r\n" +
		"\r\n" +
		"scheduled mail\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
	}
	wg.Done()

	/*
		m := gomail.NewMessage()

		m.SetHeader("From", from)

		for _, recipient := range to {
			m.SetHeader("To", recipient)
		}

		m.SetHeader("Subject", "Go!")

		m.SetBody("text/html", "Hello <b>from</b>go</i>!")

		d := gomail.NewDialer("smtp.gmail.com", 587, from, pass)

		if err := d.DialAndSend(m); err != nil {
			log.Printf("smtp error: %s", err)
			//wg.Done()
			//	log.Println("END WITH ERROR")
			//	return false, sharedD.New("Error sending the email")

		}
		wg.Done()
		//log.Println("END")
		//return true, nil
	*/
}
