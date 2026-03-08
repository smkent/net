package main

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"smkent.net/internal/handlers"
)

func testMux(t *testing.T) *http.ServeMux {
	t.Helper()
	_, filename, _, _ := runtime.Caller(0)
	staticDir := filepath.Join(filepath.Dir(filename), "../../static")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /smkent", handlers.Gallery)
	mux.HandleFunc("GET /smkent/", handlers.Gallery)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	return mux
}

func TestStaticImageOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/static/smkent.jpg", nil)
	w := httptest.NewRecorder()

	testMux(t).ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "image/jpeg" {
		t.Errorf("expected Content-Type image/jpeg, got %s", w.Header().Get("Content-Type"))
	}
	if w.Body.Len() == 0 {
		t.Error("expected non-empty response body")
	}
}
