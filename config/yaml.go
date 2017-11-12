package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Parameter describe a YAML config file.
type Parameter struct {
	Host     string              `yaml:"Host,omitempty"`
	Port     int                 `yaml:"Port,omitempty"`
	Thread   int                 `yaml:"Thread,omitempty"`
	Loop     int                 `yaml:"Loop,omitempty"`
	Include  []string            `yaml:"Include,omitempty"`
	Requests []map[string]string `yaml:"Requests,omitempty"`
}

// Read YAML file.
func Read(filename string) (*Parameter, error) {
	var parameter Parameter
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(source, &parameter)
	if err != nil {
		return nil, err
	}
	return &parameter, nil
}

// Yaml main function.
func Yaml(filename string) (*Parameter, error) {
	contents, err := Read(filename)
	if err != nil {
		return nil, err
	}

	if len(contents.Include) > 0 {
		for _, ymlfile := range contents.Include {
			conten, err := Read(ymlfile)
			if err != nil {
				log.Fatal(err)
			}

			for _, m := range conten.Requests {
				contents.Requests = append(contents.Requests, m)
			}
		}
	}

	return contents, nil
}
