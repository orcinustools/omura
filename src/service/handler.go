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
	"github.com/Jeffail/gabs"
)

// GETIndex is root endpoint for get System info
func GETIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	path := os.Getenv("GOPATH") + "/bin/repository/"
	products := gabs.New()

	ctgDir, _ := ioutil.ReadDir(path)
    for _, ctg := range ctgDir {
    	if ctg.IsDir() == true && ctg.Name() != ".git"{
    		//CtgName = append(CtgName,f.Name())
    		pdcPath := path + ctg.Name() + "/"
    		pdcDir, _ := ioutil.ReadDir(pdcPath)
    		for _, pdc := range pdcDir {
    			manifestPath := pdcPath + pdc.Name() + "/manifest.json"
    			manifestFile, err := ioutil.ReadFile(manifestPath)
				if err != nil {
					log.Printf("manifestFile.Get err   #%v ", err)
				}
				manifestData, err := gabs.ParseJSON([]byte(manifestFile))
    			products.SetP(manifestData.Path("name").Data().(string), "categories."+ctg.Name()+"."+pdc.Name()+".name")
    			products.SetP(manifestData.Path("title").Data().(string), "categories."+ctg.Name()+"."+pdc.Name()+".title")
    			products.SetP("statics/"+ ctg.Name() + "/" + pdc.Name() + "/" +manifestData.Path("logo").Data().(string), "categories."+ctg.Name()+"."+pdc.Name()+".logo")
    			products.SetP(manifestData.Path("description").Data().(string), "categories."+ctg.Name()+"."+pdc.Name()+".description")
    		}
    	}
    }

	fmt.Println(products.String())
	w.Write([]byte(products.String()))
	return
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

	items := []Product{}
	response := ResponseFormat{items}

	// parser manifest.json

	manifestPath := path + "manifest.json"
	manifestFile, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		log.Printf("manifestFile.Get err   #%v ", err)
	}
	manifestData, err := gabs.ParseJSON([]byte(manifestFile))

	logo, _ := manifestData.Path("logo").Data().(string)
	title, _ := manifestData.Path("title").Data().(string)

	//add manifest

	mnfsPrp := Manifest{
					Logo: "statics/"+ps.ByName("category") + "/" + ps.ByName("product") + "/" + logo,
					Title: title,		
				}
    mnfsData := Product{
    	Manifest: mnfsPrp,
    	Service: yamlData["services"].Service,
    }

    //mnfsJson, err := json.Marshal(mnfsData)
    //fmt.Println(string(mnfsJson))

    //add data to items
	response.AddItem(mnfsData)

	children, _ := manifestData.S("dependencies").Children()
	for _, child := range children {
		//fmt.Println(child.Data().(string))

		var yamlData map[string]Product

		yamlFile, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/bin/repository/" + child.Data().(string) + "/orcinus.yml")
		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
		}
		if err := yaml.Unmarshal([]byte(yamlFile), &yamlData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		delete(yamlData, "stack")

		ctg := os.Getenv("GOPATH") + "/bin/repository/" + child.Data().(string) + "/"
		manifestPath := ctg + "manifest.json"
		manifestFile, err := ioutil.ReadFile(manifestPath)
		if err != nil {
			log.Printf("manifestFile.Get err   #%v ", err)
		}
		manifestData, err := gabs.ParseJSON([]byte(manifestFile))

		logo, _ := manifestData.Path("logo").Data().(string)
		title, _ := manifestData.Path("title").Data().(string)

		//add manifest

		mnfsPrp := Manifest{
						Logo: "statics/"+child.Data().(string)+"/"+logo,
						Title: title,		
					}
	    mnfsData := Product{
	    	Manifest: mnfsPrp,
	    	Service: yamlData["services"].Service,
	    }

		response.AddItem(mnfsData)
	}


	resps, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Println(string(resps))
	w.Write(resps)
	return
}
