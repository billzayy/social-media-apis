package utils

import (
	"log"
	"net/smtp"
)

func SendMail(from string, to []string) {
	// SMTP server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Sender data.
	// from := "your-email@gmail.com"
	password := "cfch rdyn ahql jnwd" // Use App Password if using Gmail

	// Receiver email address.
	// to := []string{"recipient@example.com"}

	// Email message.
	message := []byte("Subject: Hello from Go!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal("Error sending email:", err)
	}

	log.Println("Email sent successfully.")
}
