package main

import (
	"log"
	"net/http"
	Handlers "github.com/fseda/url-shortener/src/handlers"
	Config "github.com/fseda/url-shortener/src/config"
)

func main() {
	Config.InitializeDB()

	http.HandleFunc("/", Handlers.IndexHandler)
	http.HandleFunc("/view/", Handlers.ViewHandler)

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("static/styles/"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("static/scripts/"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
