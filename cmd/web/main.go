package main

import (
	"context"

	"budget-track.jpech.dev/internal/repository"
	"budget-track.jpech.dev/internal/store"
	"github.com/go-playground/form/v4"
	"github.com/jackc/pgx/v5"

	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
	transactions  store.TransactionModelInterface
	formDecoder   *form.Decoder
	repo          repository.Queries
}

const (
	dbHost     = "db"
	dbPort     = 5432
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

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer conn.Close(ctx)

	repo := repository.New(conn)

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	app := &application{
		logger:        logger,
		templateCache: templateCache,
		formDecoder:   formDecoder,
		repo:          *repo,
	}

	logger.Info("starting server", "port", *port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
