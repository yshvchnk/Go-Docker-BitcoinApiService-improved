package main

import (
	"bitcoin-app/env"
	"bitcoin-app/handler"
	"bitcoin-app/store"
	rate "bitcoin-app/service/rate"
	send "bitcoin-app/service/send"
	subscribe "bitcoin-app/service/subscribe"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
)

const storagePath = "../emails.json"

func main() {
	env.LoadEnv()

	port := os.Getenv("PORT")

	router := chi.NewRouter()

	currencyAPI := rate.NewCurrencyAPIProvider()

	currencyRateHandler := handler.NewCurrencyRateHandler(currencyAPI)

	emailStorage, err := store.NewEmailStorage(storagePath)
	if err != nil {
		log.Fatal(err)
	}

	emailSender := &send.EmailSenderPath{
		StoragePath: storagePath,
	}

	emailHandler, err := handler.NewEmailSendHandler(*emailStorage, emailSender)
	if err != nil {
		log.Fatal(err)
	}

	emailServiceSubscribe := subscribe.NewEmailServiceSubscribe()

	router.Get("/api/rate", currencyRateHandler.HandleCurrencyRate)
	router.Post("/api/subscribe", func(w http.ResponseWriter, r *http.Request) {
		handler.HandleSubscribeEmails(w, r, emailServiceSubscribe)
	})
	router.Post("/api/sendEmails", emailHandler.HandleSendEmails)

	log.Println("Server started on port", port)

	serverErr := http.ListenAndServe(":"+port, router)
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}