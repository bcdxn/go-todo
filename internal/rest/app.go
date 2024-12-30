package rest

import (
	"log/slog"
	"net/http"

	"github.com/bcdxn/go-todo/internal/rest/middleware"
	"github.com/bcdxn/go-todo/internal/store/model"
)

func NewApp(
	logger *slog.Logger,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, logger, model.ToDoInMemory{})
	// Add global middlewares
	var handler http.Handler = mux
	// Note - middleware is executed in reverse order that it declared
	handler = middleware.RequestLogger(logger)(handler)
	handler = middleware.RequestID(handler, logger)
	return handler
}
