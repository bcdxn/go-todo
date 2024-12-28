package rest

import (
	"log/slog"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
) {
	mux.Handle("/api/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test complete"))
	}))
}
