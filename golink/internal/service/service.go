package service

import (
	"golink/internal/domain"
	"golink/internal/infra/env"
	"golink/internal/infra/errapp"
	"golink/internal/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type ServiceInterface interface {
	CreateUserAdmin(userRequest domain.UserRequest) (*domain.UserResponse, *errapp.ErrResponse)
}

type Service struct {
	repo     repository.RepositoryInterface
	env      *env.ENV
	validate *validator.Validate
}

func NewService(repo repository.RepositoryInterface, env *env.ENV) *Service {
	return &Service{repo: repo, env: env, validate: validator.New()}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
