package rest

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/bcdxn/go-todo/internal/store/repository"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	toDoRepository repository.ToDo,
) {
	mux.Handle("GET /api/v1/todos", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todos, err := toDoRepository.All()
		if err != nil {
			logger.ErrorContext(r.Context(), "unable to retrieve todos", "err", err)
		}

		var todosRes []toDo
		for _, todo := range todos {
			todosRes = append(todosRes, restToDoFromDomain(todo))
		}
		res, err := json.Marshal(todosRes)
		if err != nil {
			logger.ErrorContext(r.Context(), "unable to marshal todos", "err", err)
		}

		w.Write(res)
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
