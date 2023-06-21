package handlers

import (
	"bitcoin-app/files"
	"fmt"
	"net/http"
)

func HandleSubscribe(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")

	err = subscribeEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Email added")
}

func subscribeEmail(email string) error {

	subscribed, err := files.IsEmailSubscribed(email)
	if err != nil {
		return fmt.Errorf("failed to check subscription: %v", err)
	}

	if subscribed {
		return fmt.Errorf("email already subscribed")
	}

	err = files.SaveEmailToFile(email)
	if err != nil {
		return fmt.Errorf("failed to save email: %v", err)
	}

	return nil
}