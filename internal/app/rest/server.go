package rest

import (
	"log/slog"
	"net/http"

	"github.com/bcdxn/go-todo/internal/app/rest/middleware"
	"github.com/bcdxn/go-todo/internal/todo"
)

func NewServer(
	logger *slog.Logger,
	todoRepository todo.Repository,
) http.Handler {
	// Instantiate services
	todoService := todo.Service{Repository: todoRepository}
	// Create REST API router
	mux := http.NewServeMux()
	addRoutes(mux, logger, todoService)
	// Add global middlewares
	var handler http.Handler = mux
	// Note - middleware is executed in reverse order that it declared
	handler = middleware.RequestLogger(logger)(handler)
	handler = middleware.RequestID(handler, logger)
	return handler
}
