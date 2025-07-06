package utils

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(to string, username string, verificationLink string) error {
	mailSettings := gomail.NewMessage()
	mailSettings.SetHeader("From", os.Getenv("EMAIL_FROM"))
	mailSettings.SetHeader("To", to)
	mailSettings.SetHeader("Subject", "Email Verification")
	mailSettings.SetBody("text/html", `
		<h1>Welcome to ELEVATE INVIX</h1>
		<p>Hi `+username+`,</p>
		<p>Thank you for registering with us. Please click the link below to verify your email address:</p>		
		<p><a href="`+verificationLink+`">Verify Email</a></p>
		<p>If you did not register, please ignore this email.</p>
		<p>Best regards,</p>
		<p>ELEVATE INVIX Team</p>
	`)
	
	dialer := gomail.NewDialer(
		os.Getenv("EMAIL_HOST"),
		getEnvAsInt("EMAIL_PORT", 587),
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)
	if err := dialer.DialAndSend(mailSettings); err != nil {
		return err
	}
	fmt.Println("Verification email sent to: ", to)
	return nil
}

func getEnvAsInt(name string, defaultVal int) int {
	valStr := os.Getenv(name)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}
