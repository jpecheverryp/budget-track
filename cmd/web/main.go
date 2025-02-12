package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", showIndex)
	mux.HandleFunc("GET /login", showLogin)

	log.Printf("starting server on :%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	log.Fatal(err)
}
