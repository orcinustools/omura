package main

import (
	"log"
	"net/http"
	"omura/src/service"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	directory := os.Getenv("GOPATH") + "/bin/repository/"
	router := httprouter.New()
	router.GET("/apis", service.GETIndex)
	router.GET("/apis/:category", service.GETCategory)
	router.GET("/apis/:category/:product", service.GETProduct)
	router.ServeFiles("/statics/*filepath", http.Dir(directory))

	log.Printf("Serving %s on HTTP port: %s\n", directory, ":8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
