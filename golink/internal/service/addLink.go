package service

import (
	"errors"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"uuid"
)

func (s *Service) AddLinkAdmin(userID uuid.UUID, linkResquest domain.LinkRequest) (*domain.UserResponse, *errapp.ErrResponse) {
	user := domain.User{}
	if err := s.repo.FindByID(userID, &user); err != nil {
		if errors.Is(err, errapp.ErrUserNotFound) {
			return nil, errapp.NewBadRequestError("Nenhum usuário encontrado!", nil)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	if !user.Admin {
		return nil, errapp.NewUnauthorizedError("Não autorizado", nil)
	}

	if len(user.Links) >= 2 {
		return nil, errapp.NewBadRequestError("O Máximo de links foi alcançado", nil)
	}

	link := domain.Link{
		ID:    uuid.New(),
		Title: linkResquest.Title,
		Link:  linkResquest.Link,
		Count: 0,
	}

	err := s.repo.PushLink(userID, link)
	if err != nil {
		if errors.Is(err, errapp.ErrUserNotFound) {
			return nil, errapp.NewBadRequestError("Nenhum usuário encontrado!", nil)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	if err := s.repo.FindByID(userID, &user); err != nil {
		if errors.Is(err, errapp.ErrUserNotFound) {
			return nil, errapp.NewBadRequestError("Nenhum usuário encontrado!", nil)
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
