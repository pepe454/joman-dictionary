package db

import (
	"context"
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	ctx := context.Background()
	pool, err := Connect(ctx)
	if err != nil {
		t.Fatalf("Connect() error: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		t.Fatalf("Ping() after connect error: %v", err)
	}

	// There should be 3 languages in the dictionary.language table as per the test data setup.
	rows, err := pool.Query(ctx, "SELECT language, language_alt FROM dictionary.language")
	if err != nil {
		t.Fatalf("Query languages error: %v", err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var language string
		var languageAlt *string
		if err := rows.Scan(&language, &languageAlt); err != nil {
			t.Fatalf("Scan error: %v", err)
		}
		if languageAlt != nil {
			fmt.Printf("language=%s, language_alt=%s\n", language, *languageAlt)
		} else {
			fmt.Printf("language=%s, language_alt=<nil>\n", language)
		}
		count++
	}

	if err := rows.Err(); err != nil {
		t.Fatalf("rows error: %v", err)
	}

	if count != 3 {
		t.Errorf("expected 3 languages, got %d", count)
	}
}
