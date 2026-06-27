package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/pepe454/joman-dictionary/internal/api"
	"github.com/pepe454/joman-dictionary/internal/db"
	"github.com/pepe454/joman-dictionary/internal/repository"
)

func main() {
	ctx := context.Background()

	// Connect to the database
	pool, err := db.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer pool.Close()

	// Create the query handler and wire up routes
	q := repository.New(pool)
	h := api.NewAPIHandler(q)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", h.Hello)
	mux.HandleFunc("GET /categories", h.GetCategories)
	mux.HandleFunc("GET /categories/{id}/translations", h.GetTranslationsForCategory)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
