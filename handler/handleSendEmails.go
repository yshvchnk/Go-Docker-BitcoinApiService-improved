package handler

import (
	"bitcoin-app/store"
	"bitcoin-app/service"
	"fmt"
	"net/http"
	"github.com/pkg/errors"
)

type EmailHandler struct {
	EmailStorage store.EmailStorage
	BitcoinRate  BitcoinRateProvider
	EmailService *service.EmailSenderDetails
}

type BitcoinRateProvider interface {
	GetBitcoinRate() (float64, error)
}

func NewEmailHandler(storagePath string, rateProvider BitcoinRateProvider) (*EmailHandler, error) {
	storage, err := store.NewEmailStorage(storagePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create email storage")
	}

	emailSender := service.NewEmailSenderDetails(storage)

	handler := &EmailHandler{
		EmailStorage: *storage,
		BitcoinRate:  rateProvider,
		EmailService: emailSender,
	}

	return handler, nil
}

func (h *EmailHandler) HandleSendEmails(w http.ResponseWriter, r *http.Request) {
	rate, err := h.BitcoinRate.GetBitcoinRate()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	emails, err := h.EmailStorage.GetEmailsFromFile()
	if err != nil {
		err := errors.Wrap(err, "Failed to load email addresses")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success := h.EmailService.SendEmails(emails, rate)

	if !success {
		errMsg := fmt.Sprintf("Failed to send %d emails", len(emails))
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}