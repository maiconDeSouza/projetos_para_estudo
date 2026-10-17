package handler

import (
	"encoding/json"
	"fmt"
	"golink/internal/domain"
	"golink/internal/infra/env"
	"golink/internal/infra/errapp"
	"net/http"
	"uuid"
)

func (h *Handler) CreateUserAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idString := r.Context().Value(env.UserSubKey).(string)
	id, err := uuid.Parse(idString)
	if err != nil {
		errRes := errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	userRequest := domain.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		errRes := errapp.NewBadRequestError("Erro no json: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	userResponse, errRes := h.service.CreateUserAdmin(id, userRequest)
	if errRes != nil {
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userRequest := domain.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		errRes := errapp.NewBadRequestError("Erro no json: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	userResponse, errRes := h.service.CreateUser(userRequest)
	if errRes != nil {
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idString := r.Context().Value(env.UserSubKey)
	fmt.Println(idString)
	// _, err := uuid.Parse(idString)
	// if err != nil {
	// 	errResp := errapp.NewBadRequestError("Uuid no formato errado: "+err.Error(), err)
	// 	w.WriteHeader(errResp.Code)
	// 	json.NewEncoder(w).Encode(errResp.Code)
	// 	return
	// }

	isAdminString := r.Context().Value(env.UserAdminKey)
	isAdmin := isAdminString.(bool)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"Admin": isAdmin,
		"id":    idString,
	})
}
