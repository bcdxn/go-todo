package restapi

import (
	"net/http"

	"github.com/bcdxn/go-todo/pkg/restapi/json_coder"
	"github.com/bcdxn/go-todo/pkg/restapi/services"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
)

func MakeHandlerToDosGET(logger hclog.Logger, s services.GetToDoService) httprouter.Handle {
	todos := s.GetToDos()

	var res []ToDoResponse = make([]ToDoResponse, 0, len(todos))

	for _, todo := range todos {
		res = append(res, getResponseToDoFromServiceToDo(todo))
	}

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		logger.Trace("handler", "func", "HandlerToDosGet")

		err := json_coder.Encode(w, r, 200, res)
		if err != nil {
			logger.Error("error occurred writing http response", err)
			http.Error(w, "something went wrong!", 500)
		}
	}
}

type ToDoResponse struct {
	Id     string `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}

func getResponseToDoFromServiceToDo(t services.ToDo) ToDoResponse {
	return ToDoResponse{
		Id:     t.Id,
		Task:   t.Task,
		IsDone: t.IsDone,
	}
}
