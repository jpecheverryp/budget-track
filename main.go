package main

import (
	"fmt"
	"log"
	"net/http"
)

type config struct {
	port int
}

func NewConfig() *config {
	return &config{
		port: 8080,
	}
}

type application struct {
	config config
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.config.port),
	}
	log.Println("Starting Server", "addr", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	config := NewConfig()

	app := &application{
		config: *config,
	}

	err := app.serve()
	if err != nil {
		log.Fatal(err)
	}
}
