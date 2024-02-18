package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed stylesheets
//go:embed bundle.js
//go:embed index.html
var web embed.FS

func webServer() {
	port := "8080"
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.FS(web)))
	log.Printf("Starting web %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), router))
}

func main() {
	webServer()
}
