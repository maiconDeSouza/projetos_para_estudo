package errapp

import (
	"errors"
	"net/http"
)

var (
	ErrEmailAlreadyInUse = errors.New("E-mail em uso")
	ErrUserNotFound      = errors.New("Usuário não encontrado")
	ErrNoDocuments       = errors.New("Pagina não encontrada")
)

type ErrResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Err     error  `json:"error"`
}

func NewErrResponse(msg string, code int, err error) *ErrResponse {
	return &ErrResponse{Message: msg, Code: code, Err: err}
}

func NewBadRequestError(msg string, err error) *ErrResponse {
	return &ErrResponse{Message: msg, Err: err, Code: http.StatusBadRequest}
}

func NewInternalServerError(msg string, err error) *ErrResponse {
	return &ErrResponse{Message: msg, Err: err, Code: http.StatusInternalServerError}
}

func NewNotFoundError(msg string, err error) *ErrResponse {
	return &ErrResponse{Message: msg, Err: err, Code: http.StatusNotFound}
}

func NewForbiddenrError(msg string, err error) *ErrResponse {
	return &ErrResponse{Message: msg, Err: err, Code: http.StatusForbidden}
}

func NewUnauthorizedError(msg string, err error) *ErrResponse {
	return &ErrResponse{Message: msg, Err: err, Code: http.StatusUnauthorized}
}
