package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// RequestId is a type used as the key to fetch the request id from the request context
type RequestId struct{}

type RootRequestIdMiddleware struct {
	h http.Handler
}

func (m *RootRequestIdMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), RequestId{}, uuid.New())
	m.h.ServeHTTP(w, r.WithContext(ctx))
}

func NewRootRequestIdMiddleware(handlerToWrap http.Handler) *RootRequestIdMiddleware {
	return &RootRequestIdMiddleware{
		h: handlerToWrap,
	}
}
