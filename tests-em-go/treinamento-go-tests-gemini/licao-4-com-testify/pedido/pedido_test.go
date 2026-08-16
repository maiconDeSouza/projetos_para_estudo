package pedido

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPagamento struct {
	mock.Mock
}

func (m *MockPagamento) Pagar(valor float64) bool {
	mockPag := m.Called(valor)

	return mockPag.Bool(0)
}

func TestFecharPedido(t *testing.T) {
	t.Run("Sucesso!", func(t *testing.T) {
		mock := new(MockPagamento)
		mock.On("Pagar", 23.00).Return(true)

		p := ProcessadorPedido{Provedor: mock}

		err := p.FecharPedido(23.00)

		assert.NoError(t, err)

		mock.AssertExpectations(t)
	})
}
