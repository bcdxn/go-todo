package todo

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func AllToDosHandlerFunc(logger *slog.Logger, service Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// retrieve todos from data store
		todos, err := service.All()
		if err != nil {
			logger.ErrorContext(r.Context(), "unable to retrieve todos", "err", err)
		}
		// Convert from service type to response type
		var resBody []responseToDo
		for _, todo := range todos {
			resBody = append(resBody, responseToDo{
				ID:     todo.ID,
				Task:   todo.Task,
				IsDone: todo.IsDone,
			})
		}
		// Marshal the response body
		res, err := json.Marshal(resBody)
		if err != nil {
			logger.ErrorContext(r.Context(), "unable to marshal todos", "err", err)
		}
		// Write the response
		w.Write(res)
	})
}

type responseToDo struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}
