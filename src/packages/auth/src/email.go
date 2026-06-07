package src

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(toEmail string, Otp string) error {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	if from == "" || password == ""{
		return fmt.Errorf("SMTP credentials are invalid")
	}

	auth := smtp.PlainAuth(
		"",
		from,
		password,
		"smtp.gmail.com",
	)

	subject := "Subject: Verification OTP"
	body := fmt.Sprintf(
		"Your password reset OTP is %s\r\nThis otp expires in 2 miniutes.",
		Otp,
	)

	message := []byte(subject + "\r\n" + body)

	return smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		from,
		[]string{toEmail},
		message,
	)
}