package main

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"smkent.net/internal/handlers"
)

func testHandler(t *testing.T) http.Handler {
	t.Helper()
	_, filename, _, _ := runtime.Caller(0)
	staticDir := filepath.Join(filepath.Dir(filename), "../../static")
	return handlers.New(handlers.TemplateFS).Handler(staticDir)
}

func TestStaticImageOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/static/smkent.jpg", nil)
	w := httptest.NewRecorder()

	testHandler(t).ServeHTTP(w, req)

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
