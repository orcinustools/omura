package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/orcinustools/omura/src/service"
)

func main() {
	router := httprouter.New()
	router.GET("/", service.GETIndex)
	router.GET("/:category", service.GETCategory)
	router.GET("/:category/:product", service.GETProduct)

	log.Fatal(http.ListenAndServe(":8080", router))
}
