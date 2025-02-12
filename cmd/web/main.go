package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", getIndex)
    mux.HandleFunc("GET /transactions/view/{id}", getTransactionView)
    mux.HandleFunc("GET /transactions/create", getTransactionCreate)
    mux.HandleFunc("POST /transactions/create", postTransactionCreate)


	log.Printf("starting server on :%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	log.Fatal(err)
}
