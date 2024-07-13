package restapi

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

func handlerSomething(logger hclog.Logger) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			logger.Debug("neato")
			fmt.Fprintf(writer, "Hello World, %s", request.URL.Path[1:])
		},
	)
}

func addRoutes(
	mux *http.ServeMux,
	logger hclog.Logger,
) {
	mux.Handle("/", handlerSomething(logger))
}
