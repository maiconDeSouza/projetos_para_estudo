package email

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockEmail struct {
	mock.Mock
}

func (m *MockEmail) Enviar(destinatario string, mensagem string) error {
	args := m.Called(destinatario, mensagem)

	return args.Error(0)
}

func TestEnviarBoasVindas(t *testing.T) {
	mockEmail := new(MockEmail)

	mockEmail.On("Enviar", "user@teste.com").Return(nil)
}
