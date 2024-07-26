package rest

import (
	"github.com/bcdxn/go-todo/pkg/rest/handlers"
	"github.com/bcdxn/go-todo/pkg/todo"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
)

func addRoutes(
	router *httprouter.Router,
	logger hclog.Logger,
	todoService todo.Service,
) {
	router.GET("/api/todos", handlers.HandlerToDosGet(logger, todoService))
	router.GET("/api/todos/:id", handlers.HandlerToDosGetById(logger, todoService))
}
