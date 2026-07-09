package routes

import (
	"api-sistema/internal/handlers"
	"net/http"
)

func InitRoutes(handlers *handlers.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.HomeTest)

	return mux
}
