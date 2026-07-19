package handlers

import (
	"httpgordle/internal/api"
	"httpgordle/internal/handlers/getstatus"
	"httpgordle/internal/handlers/newgame"
	"net/http"
)

// NewRouter returns a router that listens for requests
// to the following endpoints:
//   - Create a new game;
//
// The provided router is ready to serve.
func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoutes, newgame.Handle)
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handle)

	return r
}

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(api.NewGameRoutes, newgame.Handle)

	return mux
}
