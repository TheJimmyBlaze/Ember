package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/thejimmyblaze/ember/common"
)

type Router interface {
	MethodFunc(method, pattern string, h http.HandlerFunc)
}

type RouteHandler interface {
	Route(r Router)
}

type APIHandler struct {
	authority common.Authority
}

func Start(authority common.Authority) error {

	log.Print("Starting Ember CA API server...")
	defer authority.Shutdown()

	router := chi.NewRouter()
	api := &APIHandler{
		authority: authority,
	}
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

func (h *APIHandler) Route(r Router) {

	log.Print("Registering Routes...")
	r.MethodFunc("GET", "/version", h.Version)
}
