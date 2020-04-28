package functions

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func UnmarshalFile(file string, t interface {}) {
	UnmarshalYaml(ReadFile(file), t)
}

func ReadFile(file string) *[]byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("error reading file " +file+ " with error ", err)
	}

	return &content
}

func UnmarshalYaml(content *[]byte, t interface{}) {
	err := yaml.Unmarshal(*content, t)
	if err != nil {
		log.Println("error unmarshalling to %T", t)
	}
}
