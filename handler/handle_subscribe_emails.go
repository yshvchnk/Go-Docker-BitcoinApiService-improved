package handler

import (
	service "bitcoin-app/service/subscribe"
	"errors"
	"fmt"
	"net/http"
)

func HandleSubscribeEmails(w http.ResponseWriter, r *http.Request, emailService *service.EmailServiceSubscribe) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")

	// emailServiceSubscribe := service.NewEmailServiceSubscribe()

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
	fmt.Fprintln(w, "Email have been added")
}

