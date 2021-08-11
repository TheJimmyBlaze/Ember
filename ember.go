package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/thejimmyblaze/ember/api"
	"github.com/thejimmyblaze/ember/authority"
	"github.com/thejimmyblaze/ember/config"
	"github.com/thejimmyblaze/ember/version"
)

const configFileName = "ember.json"

func main() {
	log.Printf("Ember - X.509 Crypto Service - %s", version.BuildVersion)
	log.Printf("Build Time: %s", version.BuildTime)
	log.Printf("Build Hash: %s", version.BuildHash)

	authority, err := configure(configFileName)
	if err != nil {
		log.Fatal(err)
	}

	err = start(authority)
	if err != nil {
		log.Fatal(err)
	}
}

func configure(fileName string) (*authority.Authority, error) {
	config, err := config.LoadConfiguration(fileName)
	if err != nil {
		return nil, err
	}

	authority, err := authority.New(config)
	return authority, err
}

func start(authority *authority.Authority) error {
	log.Print("Starting Ember CA API server...")

	router := chi.NewRouter()
	api := api.New(authority)
	api.Route(router)

	config := authority.Config
	host := fmt.Sprintf("%s:%d", config.Address, config.Port)

	log.Printf("Binding to: %s", host)

	log.Printf("Ember CA Started")
	err := http.ListenAndServe(host, router)
	return err
}
