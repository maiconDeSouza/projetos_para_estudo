package service

import (
	"errors"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"uuid"
)

func (s *Service) CreateUserAdmin(userRequest domain.UserRequest) (*domain.UserResponse, *errapp.ErrResponse) {
	if err := s.validate.Struct(userRequest); err != nil {
		return nil, errapp.NewBadRequestError("Problema em um dos campos: "+err.Error(), err)
	}

	if userRequest.Password != userRequest.RepeatPassword {
		return nil, errapp.NewBadRequestError("As senhas tem que ser iguais", nil)
	}

	hash, err := hashPassword(userRequest.Password)
	if err != nil {
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	user := domain.User{
		ID:       uuid.New(),
		Name:     userRequest.Name,
		Age:      userRequest.Age,
		Email:    userRequest.Email,
		Password: hash,
		Admin:    true,
		Links:    []domain.Link{},
	}

	err = s.repo.Create(&user)
	if err != nil {
		if errors.Is(err, errapp.ErrEmailAlreadyInUse) {
			return nil, errapp.NewBadRequestError("Email em uso: "+err.Error(), err)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	userResponse := domain.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
		Admin: user.Admin,
		Links: user.Links,
	}

	return &userResponse, nil
}
