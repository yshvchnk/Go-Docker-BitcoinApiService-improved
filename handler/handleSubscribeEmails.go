package handler

import (
	"bitcoin-app/service"
	"bitcoin-app/store"
	"errors"
	"fmt"
	"net/http"
)

const emailStoragePath = "../emails.json"

func HandleSubscribe(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")

	es, err := store.NewEmailStorage(emailStoragePath)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	emailService := service.NewEmailService(es)

	err = emailService.SubscribeEmail(email)
	if err != nil {
		if errors.Is(err, service.ErrEmailAlreadySubscribed) {
			http.Error(w, "Email is already subscribed", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to subscribe email", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Email added")
}

