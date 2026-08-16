package auth

import "errors"

var ErrTokenAusente = errors.New("token ausente")
var ErrFalha = errors.New("falha na conexão com serviço")
var ErrToken = errors.New("token expirado/inválido")

type ProvedorAuth interface {
	ValidarToken(token string) (valido bool, usuarioID string, err error)
}

type ValidadorSessao struct {
	Auth ProvedorAuth
}

func (v *ValidadorSessao) Autenticar(token string) (string, error) {
	if token == "" {
		return "", ErrTokenAusente
	}

	valido, usuarioID, err := v.Auth.ValidarToken(token)

	if err != nil {
		return "", err
	}

	if !valido {
		return "", ErrToken
	}

	return usuarioID, nil
}
