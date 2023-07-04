package handler

import (
	"bitcoin-app/store"
	"bitcoin-app/service/send"
	"net/http"
)

type EmailSendHandler struct {
	EmailService *service.EmailSendService
}

type EmailStorage interface {
	SendEmails() error
}

func NewEmailSendHandler(emailStorage store.EmailStorage, emailSender service.EmailSender) (*EmailSendHandler, error) {
	emailService := &service.EmailSendService{
		Storage: emailStorage,
		Sender:  emailSender,
	}

	handler := &EmailSendHandler{
		EmailService: emailService,
	}

	return handler, nil
}

func (h *EmailSendHandler) HandleSendEmails(w http.ResponseWriter, r *http.Request) {
	err := h.EmailService.SendEmails()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
