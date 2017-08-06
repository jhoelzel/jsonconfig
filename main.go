// Package jsonconfig reads a simple JSON configuration file from any path into a strongly typed struct and vice versa.
package jsonconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Read configuration from json file
// configpath = "./config/config.json"
func Read(configpath string) Configuration {

	file, err := ioutil.ReadFile(configpath)
	if err != nil {
		log.Fatalf("Could not find config file in path: %s because %s", configpath, err)
	}

	configuration := Configuration{}
	err = json.Unmarshal(file, &configuration)
	if err != nil {
		log.Fatalf("Could not Unmarshal config file in path: %s because %s", configpath, err)
	}
	return configuration
}

// configExists checks if the config file exists
// configpath = "./config/config.json"
func configExists(configpath string) bool {
	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		log.Fatalf("Could not find config file in path: %s because %s", configpath, err)
		return false
	}
	return true
}

// Write configuration to json file
// config = Configuration object
// configpath = "./config/config.json"
func Write(config Configuration, configpath string) {
	err := ioutil.WriteFile(configpath, config.toBytes(), 0644)
	if err != nil {
		log.Fatalf("Could not write config file to path: %s because %s", configpath, err)
	}
}

// toBytes converts the object to Bytes
func (conf Configuration) toBytes() []byte {
	bytes, err := json.Marshal(conf)
	if err != nil {
		log.Fatalf("Could not convert the Configuration to []byte because %s", err)
	}
	return bytes
}
