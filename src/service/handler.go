package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	yaml "gopkg.in/yaml.v2"
)

// Category struct manifest data
type Category struct {
	Name string
	OrcinusConf
}

// OrcinusConf struct manifest data
type OrcinusConf struct {
	Stack    string `json:"stack"`
	Services struct {
		Wordpress struct {
			Image       string   `json:"image"`
			Auth        bool     `json:"auth"`
			Ports       []string `json:"ports"`
			Environment []string `json:"environment"`
		} `json:"wordpress"`
	} `json:"services"`
}

func (c *OrcinusConf) getConf(category, product string) *OrcinusConf {
	yamlFile, err := ioutil.ReadFile("./repository/" + category + "/" + product + "/orcinus.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

// GETIndex is root endpoint for get System info
func GETIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// GETCategory endpoint for get categories by Name
func GETCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	category := &Category{
		Name: ps.ByName("category"),
	}

	result, err := json.Marshal(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
	return
}

// GETProduct endpoint for get product by name
func GETProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var orcinusYaml OrcinusConf
	orcinusYaml.getConf(ps.ByName("category"), ps.ByName("product"))

	result, err := json.Marshal(orcinusYaml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
	return
}
