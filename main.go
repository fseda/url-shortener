package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	tmplPath  = "./static/templates/"
	templates = template.Must(template.ParseFiles(tmplPath + "index.html"))
)

type Url struct {
	ID           int64  `json:"id"`
	ShortenedUrl string `json:"shortened_url"`
	OriginalUrl  string `json:"original_url"`
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("static/styles/"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("static/scripts/"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	var shortUrl string;
	if url == "" {
		shortUrl = ""
	} else {
		shortUrl = "bit.ly/" + strings.TrimPrefix(url, "https://")
	}

	newUrl := Url{
		ID: 1,
		ShortenedUrl: shortUrl,
		OriginalUrl: url,
	}

	renderTemplate(w, "index", newUrl)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data Url) {
	fmt.Println(data)
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
