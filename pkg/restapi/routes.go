package restapi

import (
	"github.com/bcdxn/go-todo/pkg/restapi/handlers"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
)

func addRoutes(
	mux *httprouter.Router,
	logger hclog.Logger,
) {
	// mux.Handle("/api/todos", handlers.HandlerToDosCreate(logger))
	// mux.Handle("/api/todos", handlers.HandlerToDosUpdate(logger))
	mux.GET("/api/todos", handlers.HandlerToDosGet(logger))
}
