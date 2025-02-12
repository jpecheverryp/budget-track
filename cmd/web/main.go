package main

import (
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
    templateCache map[string]*template.Template
}

func main() {
	port := flag.Int("port", 8080, "HTTP Network Port")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    templateCache, err := newTemplateCache()
    if err!=nil{
        logger.Error(err.Error())
        os.Exit(1)
    }

	app := &application{
		logger: logger,
        templateCache: templateCache,
	}

	logger.Info("starting server", "port", *port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
