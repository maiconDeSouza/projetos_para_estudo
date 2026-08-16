package usuario

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidacoes(t *testing.T) {
	t.Run("Idade invalida", func(t *testing.T) {
		err := ValidarIdade(15)
		assert.ErrorIs(t, err, ErrIdadeInvalida)
	})

	t.Run("Status ativo", func(t *testing.T) {
		status := ObterStatus(true)
		assert.Equal(t, "USUARIO_ATIVO", status)
	})
}
