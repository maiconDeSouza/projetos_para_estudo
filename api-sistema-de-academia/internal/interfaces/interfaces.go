package interfaces

import "api-sistema/internal/models"

type Database interface {
	CreateUser(user *models.User) error
}

type Service interface{}
