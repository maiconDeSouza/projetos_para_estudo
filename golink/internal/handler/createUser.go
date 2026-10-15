package handler

import (
	"encoding/json"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"net/http"
)

func (h *Handler) CreateUserAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userRequest := domain.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		errRes := errapp.NewBadRequestError("Erro no json: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	userResponse, errRes := h.service.CreateUserAdmin(userRequest)
	if errRes != nil {
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
}
