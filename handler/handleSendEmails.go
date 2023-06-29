package handler

import (
	"bitcoin-app/store"
	"fmt"
	"net/http"
	"github.com/pkg/errors"
)

type BitcoinRateProvider interface {
	GetBitcoinRate() (float64, error)
}

type BitcoinRateService struct {
	Provider BitcoinRateProvider
}

func (s *BitcoinRateService) GetRate() (float64, error) {
	return s.Provider.GetBitcoinRate()
}

type EmailSender interface {
	SendEmails(emails []string, rate float64) bool
	GetBitcoinRate() (float64, error)
}

type EmailService struct {
	Storage store.EmailStorage
	Sender  EmailSender
}

func (s *EmailService) SendEmails() error {
	emails, err := s.Storage.GetEmailsFromFile()
	if err != nil {
		return errors.Wrap(err, "Failed to load email addresses")
	}

	rate, err := s.Sender.GetBitcoinRate()
	if err != nil {
		return errors.Wrap(err, "Failed to get Bitcoin rate")
	}

	success := s.Sender.SendEmails(emails, rate)
	if !success {
		return fmt.Errorf("Failed to send %d emails", len(emails))
	}

	return nil
}

type EmailHandler struct {
	EmailService *EmailService
}

func NewEmailHandler(storagePath string, rateProvider BitcoinRateProvider,emailSender EmailSender) (*EmailHandler, error) {
	storage, err := store.NewEmailStorage(storagePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create email storage")
	}

	emailService := &EmailService{
		Storage: *storage,
		Sender:  emailSender,
	}

	handler := &EmailHandler{
		EmailService: emailService,
	}

	return handler, nil
}

func (h *EmailHandler) HandleSendEmails(w http.ResponseWriter, r *http.Request) {

	err := h.EmailService.SendEmails()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
