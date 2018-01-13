package main

import (
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}


	r := mux.NewRouter()
	r.HandleFunc("/webhook", VerificationEndPoint).Methods("GET")
	r.HandleFunc("/webhook", MessagesEndPoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":28080", r))
}
