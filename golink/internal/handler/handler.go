package handler

import (
	"golink/internal/infra/env"
	"golink/internal/service"
	"net/http"
)

type HandlerInterface interface {
	CreateUserAdmin(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.ServiceInterface
	env     *env.ENV
}

func NewHandler(service service.ServiceInterface, env *env.ENV) *Handler {
	return &Handler{service: service, env: env}
}
