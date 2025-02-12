package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	port := flag.Int("port", 8080, "HTTP Network Port")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.getIndex)
	mux.HandleFunc("GET /transactions/view/{id}", app.getTransactionView)
	mux.HandleFunc("GET /transactions/create", app.getTransactionCreate)
	mux.HandleFunc("POST /transactions/create", app.postTransactionCreate)

	logger.Info("starting server", "port", *port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
	logger.Error(err.Error())
	os.Exit(1)
}
