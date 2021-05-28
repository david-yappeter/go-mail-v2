package service

import (
	"fmt"
	"log"
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

//SendEmail Send Email
func SendEmail(recipients []string, subject string, htmlBody string) {

	configEmail := os.Getenv("CONFIG_EMAIL")
	configPassword := os.Getenv("CONFIG_PASSWORD")

	m := gomail.NewMessage()
	m.SetHeader("From", configEmail)

	addresses := make([]string, len(recipients))
	for i, recipient := range recipients {
		addresses[i] = m.FormatAddress(recipient, "")
	}

	m.SetHeader("To", addresses...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)

	configPort, _ := strconv.Atoi(os.Getenv("CONFIG_PORT"))

	d := gomail.NewDialer(os.Getenv("CONFIG_HOST"), configPort, configEmail, configPassword)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Mail sent!")
}
