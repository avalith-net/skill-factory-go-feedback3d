package services

import (
	"gopkg.in/gomail.v2"
)

//SendEmail is used to send emails
func SendEmail(email, subject, body string) bool {
	msg := gomail.NewMessage()
	msg.SetBody("text/html", body)

	msg.SetHeader("From", "vlotingaming@gmail.com")
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", subject)

	//Send the email to user
	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "vlotingaming@gmail.com", "V2H!x%CaxCeM")
	if err := d.DialAndSend(msg); err != nil {
		return false
	}
	return true
}
