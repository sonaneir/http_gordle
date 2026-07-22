package handlers

import (
	"net/http"

	"httpgordle/internal/api"
	"httpgordle/internal/handlers/getstatus"
	"httpgordle/internal/handlers/guess"
	"httpgordle/internal/handlers/newgame"
	"httpgordle/internal/repository"
)

// NewRouter returns a router that listens for requests to the following endpoints:
//   - Create a new game;
//   - Get the status of a game;
//   - Make a guess in a game.
//
// The provided router is ready to serve.
func NewRouter(db *repository.GameRepository) *http.ServeMux {
	r := http.NewServeMux()

	// Register each endpoint.
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handler(db))
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handler(db))
	r.HandleFunc(http.MethodPut+" "+api.GuessRoute, guess.Handler(db))

	return r
}
