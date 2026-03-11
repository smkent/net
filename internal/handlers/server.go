package handlers

import (
	"html/template"
	"io/fs"
	"net/http"
)

type Server struct {
	homeTemplate    *template.Template
	galleryTemplate *template.Template
}

func New(fsys fs.FS) *Server {
	return &Server{
		homeTemplate:    template.Must(template.ParseFS(fsys, "templates/base.html", "templates/index.html")),
		galleryTemplate: template.Must(template.ParseFS(fsys, "templates/base.html", "templates/gallery.html")),
	}
}

func (s *Server) Handler(staticFS fs.FS) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.Home)
	mux.HandleFunc("GET /smkent", s.Gallery)
	mux.HandleFunc("GET /smkent/", s.Gallery)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, staticFS, "favicon.ico")
	})
	mux.HandleFunc("GET /keybase.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, staticFS, "keybase.txt")
	})
	return mux
}
