package service

import (
	"errors"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
)

func (s *Service) TitleLink(title string) (*domain.Link, *errapp.ErrResponse) {
	link, err := s.repo.GetLinkByTitle(title)
	if err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			return nil, errapp.NewBadRequestError(err.Error(), nil)
		}
		return nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	return link, nil
}
