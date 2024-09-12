package utils

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(email, assunto, mensagem string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", assunto)
	m.SetBody("text/plain", mensagem)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS"))

	return d.DialAndSend(m)
}
