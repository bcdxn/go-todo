package rest

import (
	"log/slog"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	_ *slog.Logger,
) {
	mux.Handle("GET /api/v1/todos", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("list todos"))
	}))

	mux.Handle("POST /api/v1/todos", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("create todos"))
	}))

	mux.Handle("GET /api/v1/todos/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get todo - " + r.PathValue("id")))
	}))

	mux.Handle("PUT /api/v1/todos/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("put todo - " + r.PathValue("id")))
	}))

	mux.Handle("DELETE /api/v1/todos/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("delete todo - " + r.PathValue("id")))
	}))
}
