package main

import (
	"log"
	"net/http"

	"smkent.net/internal/handlers"
)

func main() {
	srv := handlers.New(handlers.TemplateFS)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", srv.Handler("static")); err != nil {
		log.Fatal(err)
	}
}
