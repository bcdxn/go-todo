package rest

import (
	"github.com/bcdxn/go-todo/pkg/rest/handlers"
	"github.com/bcdxn/go-todo/pkg/services"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
)

func addRoutes(
	router *httprouter.Router,
	logger hclog.Logger,
	services services.Services,
) {
	router.GET("/api/todos", handlers.HandlerToDosGet(logger, services.ToDo))
}
