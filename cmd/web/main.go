package main

import (
	"database/sql"

	"budget-track.jpech.dev/store"
	_ "github.com/lib/pq"

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
    transactions store.TransactionModelInterface
}

const (
	dbHost     = "db"
	dbPort     = 5432
	dbUser     = "budget-user"
	dbPassword = "budget-password"
	dbName     = "budget-track-db"
)

func main() {
	port := flag.Int("port", 8080, "HTTP Network Port")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*&psqlInfo)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:        logger,
		templateCache: templateCache,
        transactions: &store.TransactionModel{
            DB: db,
        },
	}

	logger.Info("starting server", "port", *port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dbInfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
