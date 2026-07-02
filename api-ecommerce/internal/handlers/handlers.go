package handlers

import (
	"api-ecommerce/internal/models"
	"api-ecommerce/internal/services"
	"encoding/json"
	"net/http"
)

type Handlers struct {
	services services.ServicesInterface
}

func (h *Handlers) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	prod := h.services.GetAllProducts()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prod)
}

func (h *Handlers) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	prod, err := h.services.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prod)
}

func (h *Handlers) UpProduct(w http.ResponseWriter, r *http.Request) {
	newProd := models.ProductResquest{}
	id := r.PathValue("id")

	err := json.NewDecoder(r.Body).Decode(&newProd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prod, err := h.services.UpProduct(id, &newProd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(prod)
}

func NewHandlers(services services.ServicesInterface) *Handlers {
	handlers := &Handlers{
		services: services,
	}
	return handlers
}
