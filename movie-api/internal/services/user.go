package services

import (
	"fmt"
	"movies-api/internal/models"
	"movies-api/internal/repositories"
)

type ServicesUserInterface interface {
	CreateUser(newUser models.NewUserRequest) (*models.User, error)
	Login(login models.Login) (*models.TokenJTWResponse, error)
}

type ServicesUser struct {
	repo repositories.RepoInterface
}

func NewServicesUser(repo repositories.RepoInterface) *ServicesUser {
	return &ServicesUser{repo: repo}
}

func (s *ServicesUser) CreateUser(newUser models.NewUserRequest) (*models.User, error) {
	user := models.User{}

	hash, err := hashPassword(newUser.Password)
	if err != nil {
		return nil, err
	}

	user.Email = newUser.Email
	user.Password = hash

	u, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *ServicesUser) Login(login models.Login) (*models.TokenJTWResponse, error) {
	user, err := s.repo.GetUser(login)
	if err != nil {
		return nil, err
	}

	if ok := checkPasswordHash(login.Password, user.Password); !ok {
		return nil, fmt.Errorf("Senha errrada")
	}

	token, err := generateJWT(*user)
	if err != nil {
		return nil, err
	}

	return token, nil

}
