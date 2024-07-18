package coder

import "net/http"

type Coder[T any] interface {
	Encode(w http.ResponseWriter, r *http.Request, status int, v T) error
	Decode(r *http.Request) (T, error)
}
