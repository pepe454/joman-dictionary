package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pepe454/joman-dictionary/internal/repository"
)

type APIHandler struct {
	q repository.Querier
}

func NewAPIHandler(q repository.Querier) *APIHandler {
	return &APIHandler{q: q}
}

// Hello handles the GET /hello endpoint and returns a simple greeting message.
func (h *APIHandler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

// GetCategories handles the GET /categories endpoint and returns a list of dictionary categories.
func (h *APIHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.q.ListCategories(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// GetTranslationsForCategory gets sourashtra/english translations for a given category.
func (h *APIHandler) GetTranslationsForCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid category id", http.StatusBadRequest)
		return
	}
	translations, err := h.q.TranslationsForCategory(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(translations)
}
