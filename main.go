package main

import (
	"io/ioutil"
	"log"
	"user-management/config"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	var config config.Config

	content, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
}
