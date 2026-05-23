package db_test

import (
	"context"
	"testing"

	"github.com/pepe454/joman-dictionary/internal/db"
)

func TestConnect(t *testing.T) {
	ctx := context.Background()

	pool, err := db.Connect(ctx)
	if err != nil {
		t.Fatalf("Connect() error: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		t.Fatalf("Ping() after connect error: %v", err)
	}
}
