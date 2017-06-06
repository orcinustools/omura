package service

import yaml "gopkg.in/yaml.v2"

// Initialize for initiates service package
func Initialize() (err error) {
	return
}

// Product response struct
type Product struct{
	Service map[string]Stack `json:"service"`
	Manifest Manifest `json:"manifest"`
}

// Stack response struct
type Stack struct {
	Image       string   `json:"image"`
	Auth        bool     `json:"auth"`
	Ports       []string `json:"ports"`
	Environments []string `json:"environments"`
}

type Manifest struct {
	Logo 		 string	  `json:"logo"`
	Title		 string	  `json:"title"`
}

// ResponseFormat response struct
type ResponseFormat struct {
	Stack []Product `json:"stack"`
}

// UnmarshalYAML for handling dynamic Orcinus.yml
func (e *Product) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var services map[string]Stack
	if err := unmarshal(&services); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	e.Service = services
	return nil
}

// AddItem service to ResponseFormat
func (r *ResponseFormat) AddItem(i Product) []Product {
	r.Stack = append(r.Stack, i)
	return r.Stack
}
