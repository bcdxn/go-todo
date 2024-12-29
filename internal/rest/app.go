package rest

import (
	"log/slog"
	"net/http"

	"github.com/bcdxn/go-todo/internal/rest/middleware"
)

func NewApp(
	logger *slog.Logger,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, logger)
	// Add global middlewares
	var handler http.Handler = mux
	// Note - middleware is executed in reverse order that it is applied
	handler = middleware.RequestLogger(logger)(handler)
	handler = middleware.RequestID(handler, logger)
	return handler
}
