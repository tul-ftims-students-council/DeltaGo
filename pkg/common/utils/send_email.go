package utils

import (
	"delta-go/pkg/common/config"
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendEmail(email string, subject string, body string) error {
	c, err := config.LoadConfig()

	if err != nil {
		return err
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	from := c.SMTPUsername
	password := c.SMTPPassword

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	fmt.Println("Sending email to " + email)

	d := gomail.NewPlainDialer(smtpHost, smtpPort, from, password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil
}
