package handler

import (
	"encoding/json"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	login := domain.Login{}
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		errRes := errapp.NewBadRequestError("Erro no json: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	cookie, loginResponse, errRes := h.service.Login(login)
	if errRes != nil {
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loginResponse)
}
