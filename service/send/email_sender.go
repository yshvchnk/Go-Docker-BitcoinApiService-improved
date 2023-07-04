package service

import (
	"bitcoin-app/service/rate"
	"log"
	"net/smtp"
	"os"
)

type EmailSenderPath struct {
	StoragePath string
}

type EmailSenderDetails struct {
	StoragePath string
}

func NewEmailSenderDetails(storagePath string) *EmailSenderDetails {
	return &EmailSenderDetails{
		StoragePath: storagePath,
	}
}

func (s *EmailSenderPath) SendEmails(emails []string, rate float64) bool {
	emailService := NewEmailSenderDetails(s.StoragePath)
	return emailService.SendEmails(emails, rate)
}

func (s *EmailSenderPath) GetCurrencyRate() (float64, error) {
	currencyAPI := service.NewCurrencyAPIProvider()

	rate, err := currencyAPI.GetCurrencyRate()
	if err != nil {
		return 0.0, err
	}

	return rate, nil
}

var (
	emailHost = os.Getenv("EMAIL_HOST")
	emailPort = os.Getenv("EMAIL_PORT")
	emailUsername = os.Getenv("EMAIL_USERNAME")
	emailPassword = os.Getenv("EMAIL_PASSWORD")
)

func SendEmail(emailText []byte, to string) error {
	smtpHost := emailHost
	smtpPort := emailPort
	smtpUsername := emailUsername
	smtpPassword := emailPassword

	err := smtp.SendMail(smtpHost+":"+smtpPort, smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost), emailFrom, []string{to}, emailText)
	if err != nil {
		log.Printf("Failed to send email to %s: %v\n", emailText, err)
		return err
	}

	log.Printf("Email sent to %s\n", emailText)
	return err
}

func (es *EmailSenderDetails) SendEmails(emails []string, rate float64) bool {
	success := true
	
	for _, email := range emails {
		emailText := CreateEmail(email, rate)
		err := SendEmail(emailText, email)
		if err != nil {
			log.Printf("Failed to send email to %s: %v\n", email, err)
		} else {
			success = false
		}
	}

	return success
}