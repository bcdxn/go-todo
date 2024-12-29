package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

type SLogContextHandler struct {
	slog.Handler
	Keys []any
}

func (h SLogContextHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(h.observe(ctx)...)
	return h.Handler.Handle(ctx, r)
}

func (h SLogContextHandler) observe(ctx context.Context) (as []slog.Attr) {
	for _, k := range h.Keys {
		a, ok := ctx.Value(k).(slog.Attr)
		if !ok {
			continue
		}
		a.Value = a.Value.Resolve()
		as = append(as, a)
	}
	return
}

func RequestLogger(l *slog.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// capture the current time
			start := time.Now()
			// wrap the writer with a negroni write so we can capture the response status code
			lrw := negroni.NewResponseWriter(w)
			h.ServeHTTP(lrw, r)
			// calculate duration of API invocation since 'start'
			duration := float64(time.Since(start)) / float64(time.Millisecond)
			// log the request
			l.InfoContext(
				r.Context(),
				"request",
				"status", lrw.Status(),
				"method", r.Method,
				"path", r.URL.Path,
				"duration_ms", duration,
			)
		})
	}
}
