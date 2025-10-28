package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"time"

	"github.com/anujkutal/go-drive/internal/data"
	"github.com/anujkutal/go-drive/internal/env"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type config struct {
	httpPort int
	db       struct {
		dsn string
	}
	jwt struct {
		secretKey string
	}
}

type application struct {
	config config
	logger *slog.Logger
	models data.Models
}

func main() {
	var cfg config

	cfg.httpPort = env.GetInt("HTTP_PORT", 4000)
	cfg.db.dsn = env.GetString("DB_DSN", "")
	cfg.jwt.secretKey = env.GetString("JWT_SECRET_KEY", "")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
