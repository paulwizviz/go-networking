package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	port := "3030"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to obtain working directory, %v", err)
	}
	log.Println(wd)
	wd = path.Join(wd, "cmd", "fileserver")
	http.Handle("/", http.FileServer(http.Dir(wd)))
	log.Printf("Starting file on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
