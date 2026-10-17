package handler

import (
	"encoding/json"
	"golink/internal/infra/env"
	"golink/internal/infra/errapp"
	"net/http"
	"uuid"
)

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idString := r.Context().Value(env.UserSubKey).(string)
	id, err := uuid.Parse(idString)
	if err != nil {
		errRes := errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	users, errRes := h.service.GetAllUser(id)
	if errRes != nil {
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
