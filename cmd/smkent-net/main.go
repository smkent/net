package main

import (
	"log"
	"net/http"

	"smkent.net/internal/handlers"
)

func main() {
	srv := handlers.New(handlers.TemplateFS)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", srv.Home)
	mux.HandleFunc("GET /smkent", srv.Gallery)
	mux.HandleFunc("GET /smkent/", srv.Gallery)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
