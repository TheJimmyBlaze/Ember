package main

import (
	"log"
	"net/http"

	"github.com/thejimmyblaze/ember/server"
	"github.com/thejimmyblaze/ember/version"
)

func main() {
	log.Printf("Ember - X.509 Crypto Service - %s", version.BuildVersion)
	log.Printf("Build Time: %s", version.BuildTime)
	log.Printf("Build Hash: %s", version.BuildHash)

	preconfigure()
	start()
}

func preconfigure() {
	log.Print("Beginning Preconfigure...")
	server.ConfigureRoutes()
}

func start() {
	log.Print("Starting...")
	err := http.ListenAndServe(":443", nil)
	log.Fatal(err)
}
