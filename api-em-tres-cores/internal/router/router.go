package router

import (
	"api-em-tres-cores/internal/handler"
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome a API em três cores 🇾🇪")
}

func SetupRoutes(playerHandler *handler.PlayerHandler, matchHandler *handler.MatchHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1", homeHandler)

	mux.HandleFunc("GET /api/v1/players", playerHandler.GetPlayers)
	mux.HandleFunc("POST /api/v1/players", playerHandler.CreatePlayer)
	mux.HandleFunc("PUT /api/v1/players/minutes/{id}", playerHandler.PlayingTime)

	mux.HandleFunc("GET /api/v1/matches", matchHandler.GetMatches)
	mux.HandleFunc("POST /api/v1/matches", matchHandler.CreateMatch)
	mux.HandleFunc("POST /api/v1/matches/event/{id}", matchHandler.NewEvent)
	mux.HandleFunc("PUT /api/v1/matches/{id}", matchHandler.ResultMatch)

	return mux
}
