package main

import (
	"log/slog"
	"os"
)

func main() {
	// Read configuration from environment variables
	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = ":8080"
	}

	dbDsn := os.Getenv("DB_DSN")
	if dbDsn == "" {
		dbDsn = "postgres://kotonilo:kotonilo_secret@localhost:5432/kotonilo_db?sslmode=disable"
	}

	cfg := config{
		addr: serverAddr,
		db: dbConfig{
			dsn: dbDsn,
		},
	}

	api := application{
		config: cfg,
	}
	//Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		logger.Error("Server has failed to start: %v\n", err)
		os.Exit(1)
	}
}
