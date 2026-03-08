package handlers_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"smkent.net/internal/handlers"
)

func TestOverlayFSUsesOverride(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "templates"), 0755); err != nil {
		t.Fatal(err)
	}
	want := "override content"
	if err := os.WriteFile(filepath.Join(dir, "templates", "base.html"), []byte(want), 0644); err != nil {
		t.Fatal(err)
	}

	overlay := handlers.NewOverlayFS(dir, handlers.TemplateFS)

	data, err := fs.ReadFile(overlay, "templates/base.html")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(data) != want {
		t.Errorf("expected %q, got %q", want, string(data))
	}
}

func TestOverlayFSFallsBack(t *testing.T) {
	overlay := handlers.NewOverlayFS(t.TempDir(), handlers.TemplateFS)

	embedded, err := fs.ReadFile(handlers.TemplateFS, "templates/index.html")
	if err != nil {
		t.Fatalf("unexpected error reading embedded: %v", err)
	}
	got, err := fs.ReadFile(overlay, "templates/index.html")
	if err != nil {
		t.Fatalf("unexpected error reading overlay: %v", err)
	}
	if string(got) != string(embedded) {
		t.Error("expected overlay to fall back to embedded content")
	}
}
