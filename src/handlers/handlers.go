package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fseda/url-shortener/src/models"
	"github.com/fseda/url-shortener/src/services"
	"github.com/fseda/url-shortener/src/templates"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	us := services.NewUrlService()

	id := r.URL.Query().Get("id")
	if id == "" {
		url := r.FormValue("url")
		if url == "" {
			templates.RenderTemplate(w, "index", models.Url{})
			return
		}

		baseUrl := getBaseURL(r)
		shortenedUrl, err := us.ShortenUrl(url, baseUrl)
		if err != nil {
			ErrorHandler(w, r)
		}

		templates.RenderTemplate(w, "index", models.Url{ShortenedUrl: shortenedUrl})
		return
	}

	if !isInt(&id) {
		fmt.Println("id is not int")
		ErrorHandler(w, r)
		return
	}

	intId, _ := strconv.Atoi(id)
	url, err := us.GetUrl(intId)
	if err != nil {
		fmt.Println("url not found")
		ErrorHandler(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "error", models.Url{})
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "view", models.Url{})
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := templates.ValidPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func isInt(value *string) bool {
	_, err := strconv.Atoi(*value)
	return err == nil
}

func getBaseURL(r *http.Request) string {
	proto := "http"
	if r.TLS != nil {
		proto = "https"
	}
	return fmt.Sprintf("%s://%s", proto, r.Host)
}
