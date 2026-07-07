package auth

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(toEmail string, Otp string) error {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	smtp_address := os.Getenv("SMTP_ADDRESS")

	if from == "" || password == ""{
		return fmt.Errorf("SMTP credentials are invalid")
	}

	auth := smtp.PlainAuth(
		"",
		from,
		password,
		smtp_address,
	)

	subject := "Subject: Verification OTP"
	body := fmt.Sprintf(
		"Your password reset OTP is %s\r\nThis OTP expires in 2 miniutes.",
		Otp,
	)

	message := []byte(subject + "\r\n" + body)

	return smtp.SendMail(
		smtp_address,
		auth,
		from,
		[]string{toEmail},
		message,
	)
}