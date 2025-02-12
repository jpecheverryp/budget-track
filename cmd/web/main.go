package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 8080, "HTTP Network Port")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", getIndex)
	mux.HandleFunc("GET /transactions/view/{id}", getTransactionView)
	mux.HandleFunc("GET /transactions/create", getTransactionCreate)
	mux.HandleFunc("POST /transactions/create", postTransactionCreate)

	logger.Info("starting server", "port", *port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
	logger.Error(err.Error())
	os.Exit(1)
}
