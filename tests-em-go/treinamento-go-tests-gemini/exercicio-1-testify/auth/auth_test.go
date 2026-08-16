package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuth struct {
	mock.Mock
}

func (m *MockAuth) ValidarToken(token string) (valido bool, usuarioID string, err error) {
	args := m.Called(token)

	return args.Bool(0), args.String(1), args.Error(2)
}

func TestAutenticar(t *testing.T) {
	t.Run("Token Vazio", func(t *testing.T) {
		m := new(MockAuth)

		v := ValidadorSessao{Auth: m}
		usuarioID, err := v.Autenticar("")

		assert.ErrorIs(t, err, ErrTokenAusente)

		assert.Empty(t, usuarioID)

		m.AssertExpectations(t)
	})

	t.Run("Token expirado/inválido", func(t *testing.T) {
		m := new(MockAuth)
		m.On("ValidarToken", "token_ivalido").Return(false, "", nil)

		v := ValidadorSessao{Auth: m}
		usuarioID, err := v.Autenticar("token_ivalido")

		assert.ErrorIs(t, err, ErrToken)

		assert.Empty(t, usuarioID)

		m.AssertExpectations(t)
	})

	t.Run("Autenticação com sucesso", func(t *testing.T) {
		m := new(MockAuth)
		m.On("ValidarToken", "user1").Return(true, "user1", nil)

		v := ValidadorSessao{Auth: m}
		usuarioID, err := v.Autenticar("user1")
		assert.NoError(t, err)

		assert.EqualValues(t, "user1", usuarioID)

		m.AssertExpectations(t)
	})

	t.Run("Falha", func(t *testing.T) {
		m := new(MockAuth)
		m.On("ValidarToken", "user1").Return(true, "", ErrFalha)

		v := ValidadorSessao{Auth: m}
		usuarioID, err := v.Autenticar("user1")
		assert.ErrorIs(t, err, ErrFalha)

		assert.Empty(t, usuarioID)

		m.AssertExpectations(t)
	})
}
