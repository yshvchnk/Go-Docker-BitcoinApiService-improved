package main

import (
	"bitcoin-app/handlers"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := chi.NewRouter()

	router.Get("/api/rate", handlers.HandleRate)
	router.Post("/api/subscribe", handlers.HandleSubscribe)
	router.Post("/api/sendEmails", handlers.HandleSendEmails)

	log.Println("Server started on port", port)

	serverErr := http.ListenAndServe(":"+port, router)
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}