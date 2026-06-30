package routes

import (
	"api-medical-system/internal/handlers"
	"net/http"
)

func InitRoutes(handlers *handlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Medical"))
	})

	mux.HandleFunc("POST /api/v1/medical-system/medical", handlers.CreateMedical)
	mux.HandleFunc("POST /api/v1/medical-system/medical/agenda", handlers.CreateAgenda)
	mux.HandleFunc("POST /api/v1/medical-system/medical/{crm}", handlers.GetMediacal)

	return mux
}
