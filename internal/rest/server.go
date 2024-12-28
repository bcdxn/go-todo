package rest

import (
	"log/slog"
	"net/http"
)

func NewServer(
	logger *slog.Logger,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, logger)
	return mux
}
