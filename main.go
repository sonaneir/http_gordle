package main

import (
	"log"
	"net/http"

	"httpgordle/internal/handlers"
	"httpgordle/internal/repository"
)

func main() {
	// Create the games' repository.
	db := repository.New()

	addr := ":8080"

	log.Print("Listening on ", addr, "...")

	// Start the server.
	err := http.ListenAndServe(addr, handlers.NewRouter(db))
	if err != nil {
		panic(err)
	}
}
