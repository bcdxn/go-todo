package handlers

import (
	"net/http"

	"github.com/bcdxn/go-todo/pkg/restapi/json_coder"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
)

func HandlerToDosGet(logger hclog.Logger) httprouter.Handle {
	type todo struct {
		Order int    `json:"order"`
		Task  string `json:"task"`
	}

	type response = []todo

	res := response([]todo{
		{
			Order: 0,
			Task:  "Laundry",
		},
		{
			Order: 1,
			Task:  "Dishes",
		},
	})

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		logger.Trace("handler", "func", "HandlerToDosGet")

		err := json_coder.Encode(w, r, 200, res)
		if err != nil {
			logger.Error("error occurred writing http response", err)
			http.Error(w, "something went wrong!", 500)
		}
	}
}
