package rest

import (
	"net"
	"net/http"

	"github.com/bcdxn/go-todo/pkg/config"
	"github.com/bcdxn/go-todo/pkg/services"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
)

func NewServer(
	cfg config.Config,
	logger hclog.Logger,
	services services.Services,
) *http.Server {
	// Create a new router that can handle routing parameters
	router := httprouter.New()
	// Map the entire surface of the API
	addRoutes(
		router,
		logger,
		services,
	)
	// Initialize and return the HTTP Server
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.Server.Host, cfg.Server.Port),
		Handler: router,
	}
	return httpServer
}
