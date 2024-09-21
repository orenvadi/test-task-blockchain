package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port string `json:"port"`
	Url  string `json:"url"`
}

// LoadConfig loads configuration from a JSON file or environment variables
func LoadConfig(configFile string) (*Config, error) {
	var config Config
	file, err := os.Open(configFile)
	if err != nil {
		log.Println("Could not open config file, trying to load from environment variables")
		config.Port = os.Getenv("PORT")
		config.Url = os.Getenv("URL")
		if config.Port == "" || config.Url == "" {
			return nil, err
		}
		return &config, nil
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
