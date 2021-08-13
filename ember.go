package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/thejimmyblaze/ember/api"
	"github.com/thejimmyblaze/ember/authority"
	"github.com/thejimmyblaze/ember/common"
	"github.com/thejimmyblaze/ember/config"
	"github.com/thejimmyblaze/ember/database"
	"github.com/thejimmyblaze/ember/version"
)

const configFileName = "config.json"

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

	log.Print("Configuring...")

	config, err := config.New(fileName)
	if err != nil {
		return nil, err
	}

	db, err := database.New(config)
	if err != nil {
		return nil, err
	}

	authority, err := authority.New(db, config)

	return authority, err
}

func start(authority common.Authority) error {

	log.Print("Starting Ember CA API server...")
	defer authority.Shutdown()

	router := chi.NewRouter()
	api := api.New(authority)
	api.Route(router)

	config := authority.GetConfig()
	address := config.GetAddress()
	port := config.GetPort()
	host := fmt.Sprintf("%s:%d", address, port)

	log.Printf("Binding to: %s...", host)

	log.Printf("Ember CA Started")
	err := http.ListenAndServe(host, router)
	return err
}
