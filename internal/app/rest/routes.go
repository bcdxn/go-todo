package rest

import (
	"log/slog"
	"net/http"

	"github.com/bcdxn/go-todo/internal/todo"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	toDoService todo.Service,
) {
	mux.Handle("GET /api/v1/todos", todo.AllToDosHandlerFunc(logger, toDoService))

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
