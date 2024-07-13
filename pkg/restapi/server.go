package restapi

import (
	"net/http"

	"github.com/hashicorp/go-hclog"
)

func NewServer(logger hclog.Logger, config Config) http.Handler {
	mux := http.NewServeMux()

	addRoutes(
		mux,
		logger,
	)

	var handler http.Handler = mux
	return handler
}
