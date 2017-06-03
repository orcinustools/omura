package service

import yaml "gopkg.in/yaml.v2"

// Initialize for initiates service package
func Initialize() (err error) {
	return
}

// Product response struct
type Product struct {
	Service map[string]Stack
}

// Stack response struct
type Stack struct {
	Image       string   `json:"image"`
	Auth        bool     `json:"auth"`
	Ports       []string `json:"ports"`
	Environment []string `json:"environment"`
}

// ProductManifest response struct
type ProductManifest struct {
	Name         string   `json:"name"`
	Logo         string   `json:"logo"`
	Description  string   `json:"description"`
	Dependencies []string `json:"dependencies"`
}

// ResponseFormat response struct
type ResponseFormat struct {
	Stack []Srv
}

// Srv response struct
type Srv struct {
	Name        string
	Logo        string
	Description string
	Image       string
	Auth        bool
	Ports       []string
	Environment []string
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
func (r *ResponseFormat) AddItem(i Srv) []Srv {
	r.Stack = append(r.Stack, i)
	return r.Stack
}
