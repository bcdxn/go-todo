package handlers

import (
	"net/http"

	"github.com/bcdxn/go-todo/pkg/coder"
	"github.com/bcdxn/go-todo/pkg/services"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
)

func HandlerToDosGet(logger hclog.Logger, todoService services.ToDoService) httprouter.Handle {
	res := todoService.GetToDos()

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		logger.Trace("handler", "func", "HandlerToDosGet")

		json := coder.JsonCoder[[]services.ToDo]{}
		err := json.Encode(w, r, 200, res)
		if err != nil {
			logger.Error("error occurred writing http response", err)
			http.Error(w, "something went wrong!", 500)
		}
	}
}
