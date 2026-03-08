package main

import (
	"io/fs"
	"log"
	"net/http"

	"smkent.net/internal/handlers"
)

func main() {
	staticFS, err := fs.Sub(handlers.StaticFS, "static")
	if err != nil {
		log.Fatal(err)
	}

	srv := handlers.New(handlers.TemplateFS)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", srv.Handler(staticFS)); err != nil {
		log.Fatal(err)
	}
}
