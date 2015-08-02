package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"gopkg.in/yaml.v2"
)

//Config struct for config file
type Config struct {
	AppID        string `yaml:"app_id,omitempty"`
	UserID       string `yaml:"user_id,omitempty"`
	Key          string `yaml:"key,omitempty"`
	OutputFile   string `yaml:"output_file,omitempty"`
	VaultAddress string `yaml:"vault_address,omitempty"`
}

var config Config

func init() {
	readConfig()
}

func readConfig() {
	configLocation := configLocation()
	rawConfigFile, err := ioutil.ReadFile(configLocation)
	if err != nil {
		log.Fatalf("Error reading config file %v\n", configLocation)
	}
	err = yaml.Unmarshal(rawConfigFile, &config)
	if err != nil {
		log.Fatal("Error parsing config file")
	}
}

func configLocation() string {
	configLocation := userHome() + "/.lockpick"
	if envLocation := os.Getenv("LOCKPICK_CONF"); envLocation != "" {
		configLocation = envLocation
	}
	return configLocation
}

func userHome() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Println("Error finding current user")
	}
	return currentUser.HomeDir
}
