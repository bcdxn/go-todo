package middlewares

import (
	"net/http"
	"time"

	"github.com/hashicorp/go-hclog"
)

type RootLoggingMiddleware struct {
	l hclog.Logger
	h http.Handler
}

func (m *RootLoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	m.h.ServeHTTP(w, r)
	t2 := time.Now()
	m.l.Info("",
		"method",
		r.Method,
		"url",
		r.URL.String(),
		"duration",
		t2.Sub(t1),
		"reqid",
		r.Context().Value(RequestId{}),
	)
}

func NewRootLoggingMiddleware(logger hclog.Logger, handlerToWrap http.Handler) *RootLoggingMiddleware {
	return &RootLoggingMiddleware{
		l: logger.Named("access_log"),
		h: handlerToWrap,
	}
}
