package handlers

import (
	"bitcoin-app/files"
	"bitcoin-app/mail"
	"bitcoin-app/service"
	"fmt"
	"log"
	"net/http"
	"github.com/pkg/errors"
)

func HandleSendEmails(w http.ResponseWriter, r *http.Request) {

	rate, err := service.GetBitcoinRate()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	emails, err := files.GetEmailsFromFile()
	if err != nil {
		err := errors.Wrap(err, "Failed to load email addresses")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success := sendEmails(emails, rate)

	if !success {
		errMsg := fmt.Sprintf("Failed to send %d emails", len(emails))
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sendEmails(emails []string, rate float64) bool {
	success := true

	for _, email := range emails {
		err := mail.SendEmail(email, rate)
		if err != nil {
			log.Printf("Failed to send email to %s: %v\n", email, err)
		} else {
			success = false
		}
	}

	return success
}