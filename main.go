package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fseda/url-shortener/src/config"
	"github.com/fseda/url-shortener/src/handlers"
)

func main() {
	err := config.InitializeDB()
	if err != nil {
		panic(err)
	}
	defer config.DB.Close()

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/error", handlers.ErrorHandler)
	http.HandleFunc("/view/", handlers.ViewHandler)

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("static/styles/"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("static/scripts/"))))

	port, _ := config.Config("PORT")

	fmt.Println("Listening on port " + port + "... 🔥")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
