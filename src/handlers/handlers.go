package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	custom_errors "github.com/fseda/url-shortener/src/errors"
	"github.com/fseda/url-shortener/src/models"
	"github.com/fseda/url-shortener/src/services"
	"github.com/fseda/url-shortener/src/templates"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	us := services.NewUrlService()

	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Println("id == \"\"")
		templates.RenderTemplate(w, "index", models.Url{})
		// handleError(w, custom_errors.ErrBadRequest)
		return
	}

	if !isInt(&id) {
		fmt.Println("id is not int")
		handleError(w, custom_errors.ErrBadRequest)
		return
	}

	intId, _ := strconv.Atoi(id)
	url, err := us.GetUrl(intId)
	if err != nil {
		fmt.Println("url not found")
		http.Redirect(w, r, "/error", http.StatusNotFound)
		// handleError(w, err)
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "error", models.Url{})
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {

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

	templates.RenderTemplate(w, "view", models.Url{})
}

func handleError(w http.ResponseWriter, err error) {
	switch err {
	case custom_errors.ErrNotFound, sql.ErrNoRows:
		http.Error(w, "Not Found", http.StatusNotFound)

	case custom_errors.ErrBadRequest:
		w.WriteHeader(http.StatusOK)
		// http.Error(w, "Bad Request", http.StatusBadRequest)

	case custom_errors.ErrUnauthorized:
		http.Error(w, "Unauthorized", http.StatusUnauthorized)

	default:
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func isInt(value *string) bool {
	_, err := strconv.Atoi(*value)
	return err == nil
}
