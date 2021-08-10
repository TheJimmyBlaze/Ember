package server

import (
	"log"
	"net/http"

	"github.com/thejimmyblaze/ember/endpoint"
)

func ConfigureRoutes() {
	log.Print("> Registering Routes...")
	log.Print("> Registering /info")
	http.Handle("/info", new(endpoint.InfoHandle))
}
