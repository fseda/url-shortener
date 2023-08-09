package handlers

import (
	"net/http"
	Templates "github.com/fseda/url-shortener/src/templates"
	Models "github.com/fseda/url-shortener/src/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Templates.RenderTemplate(w, "index", Models.Url{})
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	// url := r.FormValue("url")

	// var shortUrl string;
	// if url == "" {
	// 	shortUrl = ""
	// } else {
	// 	shortUrl = "bit.ly/" + ""
	// }

	// newUrl := Url{
	// 	ID: 1,
	// 	ShortenedUrl: shortUrl,
	// 	OriginalUrl: url,
	// }

	Templates.RenderTemplate(w, "view", Models.Url{})
}
