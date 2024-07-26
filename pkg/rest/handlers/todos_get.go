package handlers

import (
	"net/http"

	"github.com/bcdxn/go-todo/pkg/coder"
	"github.com/bcdxn/go-todo/pkg/todo"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
)

func HandlerToDosGet(logger hclog.Logger, todoService todo.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		res := todoService.GetToDos()
		json := coder.JsonCoder[[]todo.ToDo]{}

		err := json.Encode(w, r, 200, res)
		if err != nil {
			logger.Error("error occurred writing http response", err)
			http.Error(w, "something went wrong!", 500)
		}
	}
}

func HandlerToDosGetById(logger hclog.Logger, todoService todo.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		res := todoService.GetToDo(ps.ByName("id"))
		json := coder.JsonCoder[todo.ToDo]{}

		err := json.Encode(w, r, 200, *res)
		if err != nil {
			logger.Error("error occurred writing http response", err)
			http.Error(w, "something went wrong!", 500)
		}
	}
}
