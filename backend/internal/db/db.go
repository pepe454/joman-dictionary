// db package provides a unified singleton connection pool to the PostgreSQL database.
package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool     *pgxpool.Pool
	poolOnce sync.Once
)

// Connect initializes the singleton pool on first call and returns it.
// Subsequent calls return the existing pool without reconnecting.
func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	var connErr error
	poolOnce.Do(func() {
		cfg, err := pgxpool.ParseConfig(os.Getenv("POSTGRES_URL"))
		if err != nil {
			connErr = fmt.Errorf("unable to parse POSTGRES_URL: %w", err)
			return
		}
		pool, err = pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			connErr = fmt.Errorf("unable to create connection pool: %w", err)
		}
	})
	return pool, connErr
}
