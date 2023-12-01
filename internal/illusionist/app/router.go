package app

import (
	"github.com/f1xend/focus-grpc/internal/illusionist/controller"
	"net/http"
)

func route(c controller.Controller) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", c.Healthz)
	mux.HandleFunc("/show", c.Show)
	return mux
}
