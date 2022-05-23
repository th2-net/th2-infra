package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "port to bind server")
	path := flag.String("path", ".", "path to serve")
	flag.Parse()

	log.Printf("starting to serve path %s at port: %s", *path, *port)
	log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*path))))
}
