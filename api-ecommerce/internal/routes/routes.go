package routes

import (
	"api-ecommerce/internal/handlers"
	"net/http"
)

func InitRoutes(handlers *handlers.Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/ecommerce", handlers.GetAllProducts)
	mux.HandleFunc("GET /api/v1/ecommerce/{id}", handlers.GetProduct)
	mux.HandleFunc("PUT /api/v1/ecommerce/{id}", handlers.UpProduct)
	mux.HandleFunc("PUT /api/v1/ecommerce/{id}/amount", handlers.AddItem)

	return mux
}
