package handler

import (
	"encoding/json"
	"golink/internal/domain"
	"golink/internal/infra/env"
	"golink/internal/infra/errapp"
	"net/http"
	"uuid"
)

func (h *Handler) AddLinkAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idString := r.Context().Value(env.UserSubKey).(string)
	id, err := uuid.Parse(idString)
	if err != nil {
		errRes := errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	linkResquest := domain.LinkRequest{}
	if err := json.NewDecoder(r.Body).Decode(&linkResquest); err != nil {
		errRes := errapp.NewBadRequestError("Erro no json: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	userResponse, errRes := h.service.AddLinkAdmin(id, linkResquest)
	if errRes != nil {
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
}
