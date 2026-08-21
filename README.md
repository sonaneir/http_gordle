# http_gordle

An HTTP implementation of the word-guessing game Wordle (Gordle = Go + Wordle), built in Go with the standard library. The game logic is exposed as a small REST API: you create a game, submit letter guesses, and get back the state of the game after each attempt.

## Overview

Each game is stored server-side and identified by an ID. A player creates a new game, then guesses words one at a time. After every guess the API returns the feedback for each letter — correct position, wrong position, or absent — along with how many attempts are left and whether the game is won or lost.

## API

| Method | Path          | Description                          |
| ------ | ------------- | ------------------------------------ |
| POST   | `/games`      | Create a new game, returns a game ID |
| GET    | `/games/{id}` | Get the current state of a game      |
| PUT    | `/games/{id}` | Submit a guess for a game            |

## Project structure

```
http_gordle/
├── main.go                 # entry point: creates the repository, starts the server on :8080
└── internal/
    ├── handlers/           # HTTP layer: router and request handlers
    ├── repository/         # in-memory storage of games
    ├── gordle/             # core Wordle game logic (guessing, letter feedback)
    └── api/                # request/response models
```

The core game engine is kept separate from the HTTP layer, so the guessing logic is testable on its own and the transport can change without touching the rules.

## Running

```bash
make run
# or
go run .
```

The server listens on `:8080`.

Create a game:

```bash
curl -X POST http://localhost:8080/games
```

Submit a guess:

```bash
curl -X PUT http://localhost:8080/games/{id} \
  -H "Content-Type: application/json" \
  -d '{"guess": "hello"}'
```
