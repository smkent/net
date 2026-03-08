package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	"smkent.net/internal/handlers"
)

func main() {
	var staticFS fs.FS = handlers.StaticFS
	var templateFS fs.FS = handlers.TemplateFS

	if overridePath := os.Getenv("STATIC_OVERRIDE_PATH"); overridePath != "" {
		log.Printf("Using static overrides from %s", overridePath)
		staticFS = handlers.NewOverlayFS(overridePath, handlers.StaticFS)
		templateFS = handlers.NewOverlayFS(overridePath, handlers.TemplateFS)
	}

	staticSubFS, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}

	srv := handlers.New(templateFS)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", srv.Handler(staticSubFS)); err != nil {
		log.Fatal(err)
	}
}
