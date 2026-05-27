package test

import (
	"net/http"

	"github.com/moov-io/ach/server"
)

type Server struct {
	Repository server.Repository
	Service    server.Service
	Handler    http.Handler
}

func NewServer() *Server { _ = "STUB: not implemented"; return nil }
