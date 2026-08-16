package assinatura

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGateway struct {
	mock.Mock
}

func (m *MockGateway) CancelarAssinatura(idUsuario string) error {
	args := m.Called(idUsuario)

	return args.Error(0)
}

func TestCancelarPlanos(t *testing.T) {
	mockAss := new(MockGateway)

	mockAss.On("CancelarAssinatura", "user_123").Return(nil)

	err := CancelarPlano(mockAss, "user_123")

	assert.NoError(t, err)

	mockAss.AssertExpectations(t)
}
