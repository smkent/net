package handlers

import (
	"html/template"
	"net/http"
)

var galleryTemplate = template.Must(template.ParseFS(
	templateFS,
	"templates/base.html",
	"templates/gallery.html",
))

type image struct {
	Name   string
	Width  int
	Height int
}

type galleryData struct {
	Host   string
	Images []image
}

var images = []image{
	{"expanse", 400, 400},
	{"orville", 400, 532},
	{"galaxyquest", 400, 514},
	{"futurama", 400, 400},
	{"doctorwho", 400, 461},
	{"bsg", 400, 551},
	{"spaceball1", 200, 234},
	{"agentsmith", 329, 326},
	{"terminator", 200, 222},
	{"borg", 200, 175},
}

func Gallery(w http.ResponseWriter, r *http.Request) {
	data := galleryData{
		Host:   r.Host,
		Images: images,
	}
	if err := galleryTemplate.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
