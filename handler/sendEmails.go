package handler

import (
	"bitcoin-app/file"
	"bitcoin-app/service"
	"fmt"
	"net/http"
	"github.com/pkg/errors"
)

func HandleSendEmails(w http.ResponseWriter, r *http.Request) {

	rate, err := service.GetBitcoinRate()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	emails, err := file.GetEmailsFromFile()
	if err != nil {
		err := errors.Wrap(err, "Failed to load email addresses")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success := service.SendEmails(emails, rate)

	if !success {
		errMsg := fmt.Sprintf("Failed to send %d emails", len(emails))
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

