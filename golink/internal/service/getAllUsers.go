package service

import (
	"errors"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"uuid"
)

func (s *Service) GetAllUser(id uuid.UUID) ([]*domain.User, *errapp.ErrResponse) {
	userLogin := domain.User{}
	err := s.repo.FindByID(id, &userLogin)

	if err != nil || !userLogin.Admin {
		return nil, errapp.NewUnauthorizedError("Não autorizado", nil)
	}

	users, err := s.repo.FindAllUsers()
	if err != nil {
		if errors.Is(err, errapp.ErrUserNotFound) {
			return nil, errapp.NewBadRequestError("Nenhum usuário encontrado!", nil)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	return users, nil
}
