package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

const pgpKey = "0x0342C80999FB1A06FD0F95338392C992D92500A9"

var homeTemplate = template.Must(template.ParseFS(
	templateFS,
	"templates/base.html",
	"templates/index.html",
))

type homeData struct {
	Host   string
	PGPURL string
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := homeData{
		Host:   r.Host,
		PGPURL: fmt.Sprintf("https://keys.openpgp.org/search?q=%s", pgpKey),
	}
	if err := homeTemplate.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
