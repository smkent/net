package handlers_test

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"smkent.net/internal/handlers"
)

func testHandler(t *testing.T) http.Handler {
	t.Helper()
	staticFS, err := fs.Sub(handlers.StaticFS, "static")
	if err != nil {
		t.Fatal(err)
	}
	return handlers.New(handlers.TemplateFS).Handler(staticFS)
}

func TestHomeOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Host = "smkent.net"
	w := httptest.NewRecorder()

	testHandler(t).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "<title>smkent.net</title>") {
		t.Errorf("expected body to contain page title, got:\n%s", body)
	}
}

func TestGalleryOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/smkent", nil)
	w := httptest.NewRecorder()

	testHandler(t).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "gallery-id") {
		t.Errorf("expected body to contain gallery images, got:\n%s", body)
	}
}

func TestHomeNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
	w := httptest.NewRecorder()

	testHandler(t).ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}
}
