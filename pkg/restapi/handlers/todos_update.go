package handlers

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

func HandlerToDosUpdate(logger hclog.Logger) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			logger.Debug("neato")
			fmt.Fprintf(writer, "Update To-Do, %s", request.URL.Path[1:])
		},
	)
}
