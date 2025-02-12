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

	logger.Info("starting server", "port", *port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
