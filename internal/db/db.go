package db

import (
	"context"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
)

// New creates a new PostgreSQL connection pool
func New(addr string, maxOpenConns int, maxIdleConns int, maxIdleTime string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, err
	}

	// Set pool configuration
	// Note: pgxpool configuration works slightly differently than database/sql
	// But we can set MaxConns.
	// For now we will use the defaults from the connection string or simple defaults if needed,
	// but pgxpool.ParseConfig usually handles the connection string parameters well.

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// Connection Check
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
