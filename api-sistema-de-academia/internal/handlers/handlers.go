package handlers

import (
	"api-sistema/internal/interfaces"
	"encoding/json"
	"net/http"
)

type Handler struct {
	services interfaces.Service
}

func (h *Handler) HomeTest(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"message": "oi",
	}

	json.NewEncoder(w).Encode(res)
}

func NewHandlers(services interfaces.Service) *Handler {
	return &Handler{
		services: services,
	}
}
