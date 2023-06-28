package main

import (
	"bitcoin-app/handler"
	"bitcoin-app/service"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

const storagePath = "../emails.json"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := chi.NewRouter()

	bitcoinAPI := service.NewCoinGeckoAPI()

	bitcoinRateHandler := handler.NewBitcoinRateHandler(bitcoinAPI)

	emailHandler, err := handler.NewEmailHandler(storagePath, bitcoinAPI)
	if err != nil {
		log.Fatal(err)
	}

	router.Get("/api/rate", bitcoinRateHandler.HandleRate)
	router.Post("/api/subscribe", handler.HandleSubscribe)
	router.Post("/api/sendEmails", emailHandler.HandleSendEmails)

	log.Println("Server started on port", port)

	serverErr := http.ListenAndServe(":"+port, router)
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}