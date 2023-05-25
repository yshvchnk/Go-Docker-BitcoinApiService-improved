package main

import (
	"bitcoin-app/api"
	"log"      //logging
	"net/http" //work with http
)

func main() {
	//handlers for urls
	http.HandleFunc("/api/rate", api.HandleRate)
	http.HandleFunc("/api/subscribe", api.HandleSubscribe)
	http.HandleFunc("/api/sendEmails", api.HandleSendEmails)

	//logging
	log.Println("Server started on port 8080")

	//start server
	err := http.ListenAndServe(":8080", nil)

	//error handling
	if err != nil {
		log.Fatal(err)
	}
}