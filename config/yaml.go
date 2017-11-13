package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Parameter describe a YAML config file.
type Parameter struct {
	Host     string              `yaml:"host,omitempty"`
	Port     int                 `yaml:"port,omitempty"`
	Thread   int                 `yaml:"thread,omitempty"`
	Loop     int                 `yaml:"loop,omitempty"`
	Include  []string            `yaml:"include,omitempty"`
	Requests []map[string]string `yaml:"requests,omitempty"`
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
