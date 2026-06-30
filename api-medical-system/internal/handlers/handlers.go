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

func (h *Handlers) CreateAgenda(w http.ResponseWriter, r *http.Request) {
	medical := models.MedicalAgenda{}

	err := json.NewDecoder(r.Body).Decode(&medical)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.services.CreateAgenda(&medical)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := map[string]string{
		"message": "Agenda Criada com sucesso",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *Handlers) GetMediacal(w http.ResponseWriter, r *http.Request) {
	crm := r.PathValue("crm")

	medical, err := h.services.GetMedical(crm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(medical)
}

func (h *Handlers) DeleteMedical(w http.ResponseWriter, r *http.Request) {
	crm := r.PathValue("crm")

	medical, err := h.services.GetMedical(crm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(medical)
}

func NewHandlers(services services.ServicesInterface) *Handlers {
	return &Handlers{
		services: services,
	}
}
