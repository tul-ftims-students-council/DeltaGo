package utils

import (
	"delta-go/pkg/common/config"
	"fmt"
	"net/smtp"
)

func SendEmail(email string, subject string, body string) error {
	c, err := config.LoadConfig()

	fmt.Printf("Sending email to %s", email)

	if err != nil {
		return err
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := c.SMTPUsername
	password := c.SMTPPassword

	to := []string{
		email,
	}

	message := []byte(body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {
		return err
	}

	return nil
}
