package api

import (
	"log"
	"net/http"

	"github.com/thejimmyblaze/ember/authority"
)

type Router interface {
	MethodFunc(method, pattern string, h http.HandlerFunc)
}

type RouteHandler interface {
	Route(r Router)
}

type APIHandler struct {
	authority *authority.Authority
}

func New(authority *authority.Authority) RouteHandler {
	return &APIHandler{
		authority: authority,
	}
}

func (h *APIHandler) Route(r Router) {
	log.Print("Registering Routes...")

	r.MethodFunc("GET", "/version", h.Version)
}
