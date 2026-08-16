package usuario

import (
	"errors"
)

// Funções da aplicação ⚙️
var ErrIdadeInvalida = errors.New("idade minima e 18 anos")

func ValidarIdade(idade int) error {
	if idade < 18 {
		return ErrIdadeInvalida
	}
	return nil
}

func ObterStatus(ativo bool) string {
	if ativo {
		return "USUARIO_ATIVO"
	}
	return "USUARIO_INATIVO"
}
