package service

import (
	"errors"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
)

func (s *Service) RedirectTitleLink(title string) (*domain.Link, *errapp.ErrResponse) {
	link, err := s.repo.GetLinkByTitle(title)
	if err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			return nil, errapp.NewBadRequestError("Link não encontrado na base de dados!", nil)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	// Validação de segurança: se o link veio vazio por erro de mapeamento do BSON
	if link.Data.Link == "" {
		return nil, errapp.NewInternalServerError("Erro interno: dados do link não foram carregados corretamente", nil)
	}

	user := domain.User{}
	if err = s.repo.FindByID(link.UserID, &user); err != nil {
		if errors.Is(err, errapp.ErrUserNotFound) {
			return nil, errapp.NewBadRequestError("Nenhum usuário encontrado!", nil)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	for i := range user.Links {
		if user.Links[i].Title == title {
			user.Links[i].Count++
			break
		}
	}

	err = s.repo.UpdateUser(link.UserID, &user)
	if err != nil {
		if errors.Is(err, errapp.ErrUserNotFound) {
			return nil, errapp.NewBadRequestError("Nenhum usuário encontrado!", nil)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	linkRes := domain.Link{
		ID:        link.Data.ID,
		Title:     link.Data.Title,
		Link:      link.Data.Link,
		CreatedAt: link.Data.CreatedAt,
		UpdatedAt: link.Data.UpdatedAt,
		Count:     link.Data.Count,
	}

	return &linkRes, nil
}
