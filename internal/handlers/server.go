package handlers

import (
	"html/template"
	"io/fs"
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
