package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/segmentio/ksuid"
)

// custom key type will prevent collisions
type contextKey string

const (
	RequestIDCtxKey contextKey = "req_id"
)

func RequestID(h http.Handler, l *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), RequestIDCtxKey, slog.Attr{
			Key:   "req_id",
			Value: slog.StringValue("req_" + ksuid.New().String()),
		})
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
