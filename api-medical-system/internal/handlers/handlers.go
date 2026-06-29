package handlers

import (
	"api-medical-system/internal/models"
	"api-medical-system/internal/services"
	"encoding/json"
	"net/http"
)

type Handlers struct {
	services services.ServicesInterface
}

func (h *Handlers) CreateMedical(w http.ResponseWriter, r *http.Request) {
	medical := models.MedicalRequest{}

	err := json.NewDecoder(r.Body).Decode(&medical)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMedical, err := h.services.CreateMedical(&medical)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMedical)
}

func NewHandlers(services services.ServicesInterface) *Handlers {
	return &Handlers{
		services: services,
	}
}
