package templates

import (
	"html/template"
	"net/http"
	"sync"

	Models "github.com/fseda/url-shortener/src/models"
)

var (
	tmplPath           = "./static/templates/"
	templates          = template.Must(template.ParseGlob(tmplPath + "*.html"))
	templateCacheMutex sync.Mutex
	templateCache      = make(map[string]*template.Template)
)

func getCachedTemplate(templateName string) (*template.Template, error) {
	templateCacheMutex.Lock()
	defer templateCacheMutex.Unlock()

	tmpl, found := templateCache[templateName]
	if !found {
		tmpl, err := templates.Clone()
		if err != nil {
			return nil, err
		}

		tmpl, err = tmpl.ParseFiles(tmplPath + templateName + ".html")
		if err != nil {
			return nil, err
		}

		templateCache[templateName] = tmpl
	}

	return tmpl, nil
}

func RenderTemplate(w http.ResponseWriter, tmplName string, data Models.Url) {
	tmpl, err := getCachedTemplate(tmplName)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, tmplName+".html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
