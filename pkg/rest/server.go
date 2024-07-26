package rest

import (
	"net"
	"net/http"

	"github.com/bcdxn/go-todo/pkg/config"
	"github.com/bcdxn/go-todo/pkg/rest/middlewares"
	"github.com/bcdxn/go-todo/pkg/todo"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
)

func NewServer(
	cfg config.Config,
	logger hclog.Logger,
	todoService todo.Service,
) *http.Server {
	// Create a new router that can handle routing parameters
	router := httprouter.New()
	// Map the entire surface of the API
	addRoutes(
		router,
		logger,
		todoService,
	)
	// Apply root-level middlewares
	handler := middlewares.NewRootRequestIdMiddleware(
		middlewares.NewRootLoggingMiddleware(logger, router),
	)
	// Initialize and return the HTTP Server
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.Server.Host, cfg.Server.Port),
		Handler: handler,
	}
	return httpServer
}
