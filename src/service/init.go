package service

import yaml "gopkg.in/yaml.v2"

// Initialize for initiates service package
func Initialize() (err error) {
	return
}

// Services response struct
type Services struct {
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
	Dependencies []string `json:"dependencies"`
}

// ResponseFormat response struct
type ResponseFormat struct {
	Spec []struct {
		Name     string `json:"name"`
		Logo     string `json:"logo"`
		Services struct {
			Image       string   `json:"image"`
			Auth        bool     `json:"auth"`
			Ports       []string `json:"ports"`
			Environment []string `json:"environment"`
		} `json:"service"`
	} `json:"spec"`
}

// UnmarshalYAML for handling dynamic Orcinus.yml
func (e *Services) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var services map[string]Stack
	if err := unmarshal(&services); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	e.Service = services
	return nil
}
