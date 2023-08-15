package templates

import (
	"html/template"
	"net/http"
	"regexp"

	Models "github.com/fseda/url-shortener/src/models"
)

var (
	tmplPath  = "./static/templates/"
	templates = template.Must(template.ParseFiles(tmplPath+"index.html", tmplPath+"error.html", tmplPath+"view.html"))
	ValidPath = regexp.MustCompile("^/(|error)/([a-zA-Z0-9]+)$")
)

func RenderTemplate(w http.ResponseWriter, tmplName string, data Models.Url) {
	err := templates.ExecuteTemplate(w, tmplName+".html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

