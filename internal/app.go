package internal

import (
	"log"

	"github.com/thejimmyblaze/ember/api"
	"github.com/thejimmyblaze/ember/config"
	"github.com/thejimmyblaze/ember/pki"
	"github.com/thejimmyblaze/ember/version"
)

func Start() {

	log.Printf("Ember - X.509 Crypto Service - %s", version.BuildVersion)
	log.Printf("Build Time: %s", version.BuildTime)
	log.Printf("Build Hash: %s", version.BuildHash)

	//Create authority
	authority, err := pki.CreateAuthority(config.DefaultConfigFileName)
	if err != nil {
		log.Fatal(err)
	}

	//Start API
	err = api.Start(authority)
	if err != nil {
		log.Fatal(err)
	}
}
