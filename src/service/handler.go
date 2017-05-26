package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GETIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func GETCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Category, %s!\n", ps.ByName("category"))
}

func GETProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "%s with %s product\n", ps.ByName("category"), ps.ByName("product"))
}
