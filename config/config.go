package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pkg/errors"
)

const DefaultConfigFileName = "config.json"

type Config struct {
	Address    string `json:"address"`
	Port       int    `json:"port"`
	DBFileName string `json:"dbFileName"`
}

func New(fileName string) (*Config, error) {
	log.Printf("Loading configuration: %s...", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening config file: %s", fileName)
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing config file: %s", fileName)
	}

	config.init()
	return &config, nil
}

func (config *Config) init() {
	if config.Address == "" {
		config.Address = "0.0.0.0"
		log.Printf("Using default Address: %s", config.Address)
	}
	if config.Port == 0 {
		config.Port = 443
		log.Printf("Using default Port: %d", config.Port)
	}
	if config.DBFileName == "" {
		config.DBFileName = "ember.sqlite3"
		log.Printf("Using default DBFileName: %s", config.DBFileName)
	}
}

func (config *Config) GetAddress() string {
	return config.Address
}

func (config *Config) GetPort() int {
	return config.Port
}

func (config *Config) GetDBFileName() string {
	return config.DBFileName
}
