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
	if err := yaml.Unmarshal1([]byte(yamlFile), &yamlData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delete(yamlData, "stack")
	fmt.Printfln(yamlData);

	result, err := json.Marshal(yamlData["services"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var resultData = []byte(result)
	var resData Product
	if err := json.Unmarshal(resultData, &resData); err != nil {
		fmt.Println(err.Error())
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

	var mnf = []byte(manifestFile)
	var manif ProductManifest
	if err := json.Unmarshal(mnf, &manif); err != nil {
		fmt.Println(err.Error())
	}

	items := []Srv{}
	response := ResponseFormat{items}
	for _, v := range manifestData.Dependencies {
		deps := os.Getenv("GOPATH") + "/bin/repository/" + v + "/orcinus.yml"
		var yamlData map[string]Product
		yamlFile, err := ioutil.ReadFile(deps)
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
		var resultData = []byte(result)
		var resData Product
		if err := json.Unmarshal(resultData, &resData); err != nil {
			fmt.Println(err.Error())
		}
		for _, v := range resData.Service {
			item := Srv{
				Name:        manif.Name,
				Logo:        manif.Logo,
				Description: manif.Description,
				Image:       v.Image,
				Auth:        v.Auth,
				Ports:       v.Ports,
				Environment: v.Environment,
			}
			response.AddItem(item)
		}
	}

	resps, err := json.Marshal(response.Stack)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(resps)
	return
}
