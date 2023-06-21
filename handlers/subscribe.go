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

	subscribed, err := files.IsEmailSubscribed(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if subscribed {
		http.Error(w, "Email already subscribed", http.StatusConflict)
		return
	}

	err = files.SaveEmailToFile(email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Email added")
}