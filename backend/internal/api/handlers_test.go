package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pepe454/joman-dictionary/internal/repository"
)

// mockQuerier embeds the Querier interface so only the methods under test need implementing.
type mockQuerier struct {
	repository.Querier
	listCategories           func(ctx context.Context) ([]repository.DictionaryCategory, error)
	translationsForCategory  func(ctx context.Context, categoryID int32) ([]repository.TranslationsForCategoryRow, error)
}

func (m *mockQuerier) ListCategories(ctx context.Context) ([]repository.DictionaryCategory, error) {
	return m.listCategories(ctx)
}

func (m *mockQuerier) TranslationsForCategory(ctx context.Context, categoryID int32) ([]repository.TranslationsForCategoryRow, error) {
	return m.translationsForCategory(ctx, categoryID)
}

func TestGetCategories_Success(t *testing.T) {
	expected := []repository.DictionaryCategory{
		{CategoryID: 1, Category: "food"},
		{CategoryID: 2, Category: "body"},
	}

	h := NewAPIHandler(&mockQuerier{
		listCategories: func(ctx context.Context) ([]repository.DictionaryCategory, error) {
			return expected, nil
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	w := httptest.NewRecorder()
	h.GetCategories(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.StatusCode)
	}
	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %s", ct)
	}

	var got []repository.DictionaryCategory
	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(got) != len(expected) {
		t.Fatalf("expected %d categories, got %d", len(expected), len(got))
	}
	for i, c := range got {
		if c != expected[i] {
			t.Errorf("category[%d]: expected %+v, got %+v", i, expected[i], c)
		}
	}
}

func TestGetTranslationsForCategory_Success(t *testing.T) {
	expected := []repository.TranslationsForCategoryRow{
		{SourashtraWordID: 1, SourashtraText: "ఆహారం", TargetWordText: "food"},
	}

	h := NewAPIHandler(&mockQuerier{
		translationsForCategory: func(ctx context.Context, categoryID int32) ([]repository.TranslationsForCategoryRow, error) {
			return expected, nil
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/categories/1/translations", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()
	h.GetTranslationsForCategory(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.StatusCode)
	}
	var got []repository.TranslationsForCategoryRow
	if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(got) != len(expected) {
		t.Fatalf("expected %d rows, got %d", len(expected), len(got))
	}
}

func TestGetTranslationsForCategory_InvalidID(t *testing.T) {
	h := NewAPIHandler(&mockQuerier{})

	req := httptest.NewRequest(http.MethodGet, "/categories/abc/translations", nil)
	req.SetPathValue("id", "abc")
	w := httptest.NewRecorder()
	h.GetTranslationsForCategory(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Result().StatusCode)
	}
}

func TestGetTranslationsForCategory_DBError(t *testing.T) {
	h := NewAPIHandler(&mockQuerier{
		translationsForCategory: func(ctx context.Context, categoryID int32) ([]repository.TranslationsForCategoryRow, error) {
			return nil, errors.New("db error")
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/categories/1/translations", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()
	h.GetTranslationsForCategory(w, req)

	if w.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Result().StatusCode)
	}
}

func TestGetCategories_DBError(t *testing.T) {
	h := NewAPIHandler(&mockQuerier{
		listCategories: func(ctx context.Context) ([]repository.DictionaryCategory, error) {
			return nil, errors.New("db connection lost")
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	w := httptest.NewRecorder()
	h.GetCategories(w, req)

	if w.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Result().StatusCode)
	}
}
