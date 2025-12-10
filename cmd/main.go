package main

import (
	"log/slog"
	"os"

	"github.com/Abir-Zayn/kotoNilo/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	//Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Debug: Print current working directory
	cwd, _ := os.Getwd()
	logger.Info("Current working directory", "cwd", cwd)

	// Try loading .env file
	if err := godotenv.Load(); err != nil {
		logger.Warn("No .env file found", "error", err)
	}

	// Read configuration
	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = ":8080"
	}

	dbDsn := os.Getenv("DB_DSN")
	if dbDsn == "" {
		logger.Error("DB_DSN environment variable is required")
		os.Exit(1)
	}

	logger.Info("Starting application")

	cfg := config{
		addr: serverAddr,
		db: dbConfig{
			dsn: dbDsn,
		},
	}

	// DB Connection
	pool, err := db.New(cfg.db.dsn, 30, 30, "15m")
	if err != nil {
		logger.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	logger.Info("Connected to database")

	api := application{
		config: cfg,
		db:     pool,
	}

	if err := api.run(api.mount()); err != nil {
		logger.Error("Server has failed to start: %v\n", err)
		os.Exit(1)
	}
}
