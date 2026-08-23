package router

import (
	"api-em-tres-cores/internal/handler"
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome a API em três cores 🇾🇪")
}

func SetupRoutes(playerHandler *handler.PlayerHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1", homeHandler)
	mux.HandleFunc("GET /api/v1/players", playerHandler.GetPlayers)
	mux.HandleFunc("POST /api/v1/player", playerHandler.CreatePlayer)

	return mux
}
