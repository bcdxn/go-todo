package restapi

import (
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
)

func NewServer(logger hclog.Logger, config Config) http.Handler {
	mux := httprouter.New()

	addRoutes(
		mux,
		logger,
	)

	var handler http.Handler = mux
	return handler
}
