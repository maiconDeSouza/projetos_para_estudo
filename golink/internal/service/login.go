package service

import (
	"errors"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"net/http"
)

func (s *Service) Login(login domain.Login) (*http.Cookie, *domain.LoginResponse, *errapp.ErrResponse) {
	if err := s.validate.Struct(login); err != nil {
		return nil, nil, errapp.NewBadRequestError("Problema em um dos campos: "+err.Error(), err)
	}

	user := domain.User{}
	err := s.repo.FindByEmail(login.Email, &user)
	if err != nil {
		if errors.Is(err, errapp.ErrUserNotFound) {
			return nil, nil, errapp.NewBadRequestError("Usuário ou senha inválidos", nil)
		}
		return nil, nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	ok := checkPasswordHash(login.Password, user.Password)
	if !ok {
		return nil, nil, errapp.NewBadRequestError("Usuário ou senha inválidos", nil)
	}

	jwt, err := generateJWT(user.ID, user.Admin, s.env)
	if err != nil {
		return nil, nil, errapp.NewInternalServerError("Problema no servidor: "+err.Error(), err)
	}

	cookie := generateCookie(jwt)

	loginResponse := domain.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Admin: user.Admin,
		Links: user.Links,
	}

	return cookie, &loginResponse, nil
}
