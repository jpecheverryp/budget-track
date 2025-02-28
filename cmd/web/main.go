package main

import (
	"database/sql"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/jpecheverryp/budget-track/internal/repository"
	"github.com/jpecheverryp/budget-track/internal/service"

	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type application struct {
	logger         *slog.Logger
	service        service.Service
	sessionManager *scs.SessionManager
	repo           *repository.Queries
}

const (
	dbHost     = "localhost"
	dbPort     = 8081
	dbUser     = "budget-user"
	dbPassword = "budget-password"
	dbName     = "budget-track-db"
	sslMode    = "disable"
)

func main() {
	port := flag.Int("port", 8080, "HTTP Network Port")
	flag.Parse()
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	repo := repository.New(db)

	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		logger:         logger,
		service:        service.New(*repo),
		sessionManager: sessionManager,
		repo:           repo,
	}

	logger.Info("starting server", "port", *port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
