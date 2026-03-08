package main

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
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
