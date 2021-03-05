package config_parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var config ConfigStructure

func ParseConfig() ConfigStructure {

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Getting Current Working Directory")
	log.Printf("Acquired path: "+path)
	log.Printf("Parsing Config")
	jsonFile, err := os.Open("config.json")
	if err != nil {

		log.Fatal("Cannot open config file")

	} else {

		log.Printf("Acquired config")

	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	log.Printf("Reading config")
	json.Unmarshal(byteValue, &config)

	return config

}