package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fseda/url-shortener/src/models"
	"github.com/fseda/url-shortener/src/services"
	"github.com/fseda/url-shortener/src/templates"
)

type contextKey string

const (
	urlKey contextKey = "originalUrl"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	us := services.NewUrlService()

	url := r.FormValue("url")
	id := r.URL.Query().Get("id")
	if id == "" {
		if url == "" {
			templates.RenderTemplate(w, "index", models.Url{})
			return
		}

		// Show shortened url
		ctx := context.WithValue(r.Context(), urlKey, url)
		ViewHandler(w, r.WithContext(ctx))
		return
	}

	if !isInt(&id) {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	// Look for url with given id
	intId, _ := strconv.Atoi(id)
	url, err := us.GetUrl(intId)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	// url is found and app redirects to it
	http.Redirect(w, r, url, http.StatusFound)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "error", models.Url{})
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	us := services.NewUrlService()

	url, _ := r.Context().Value(urlKey).(string)
	fmt.Println(url)
	if url == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	baseUrl := getBaseURL(r)
	shortenedUrl, err := us.ShortenUrl(url, baseUrl)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}

	templates.RenderTemplate(w, "view", models.Url{
		OriginalUrl:  url,
		ShortenedUrl: shortenedUrl,
	})
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
