package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {

	path := "./config.yaml"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	var config SimpleProxy
	err := loadConfig(&config, path)
	if err != nil {
		log.Fatal(err)
	}

	subdomains := loadHandlers(&config)

	log.Printf("Starting server on port %d\n", config.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), subdomains)
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig(config *SimpleProxy, path string) error {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Error was encounter when trying to read %s, %v", path, err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return fmt.Errorf("Error was encounter when trying to unmarshal %s, %v", path, err)
	}

	return nil
}
