package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	yaml "gopkg.in/yaml.v2"
)

// GETIndex is root endpoint for get System info
func GETIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// GETCategory endpoint for get categories by Name
func GETCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Your "+ps.ByName("category")+" has arrived!\n")
}

// GETProduct endpoint for get product by name
func GETProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	path := os.Getenv("GOPATH") + "/bin/repository/" + ps.ByName("category") + "/" + ps.ByName("product") + "/"

	var yamlData map[string]Product
	yamlFile, err := ioutil.ReadFile(path + "orcinus.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	if err := yaml.Unmarshal([]byte(yamlFile), &yamlData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delete(yamlData, "stack")

	result, err := json.Marshal(yamlData["services"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var manifestData ProductManifest
	manifestFile, err := ioutil.ReadFile(path + "manifest.json")
	if err != nil {
		log.Printf("manifestFile.Get err   #%v ", err)
	}
	if err := json.Unmarshal([]byte(manifestFile), &manifestData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, v := range manifestData.Dependencies {
		deps := os.Getenv("GOPATH") + "/bin/repository/" + v + "/orcinus.yml"
		fmt.Printf("%v\n", deps)
	}

	w.Write(result)
	return
}
