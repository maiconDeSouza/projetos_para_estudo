package desconto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVip struct {
	mock.Mock
}

var testes = []struct {
	nome          string
	clienteID     string
	valorOriginal float64
	ehVIP         bool
	valorEsperado float64
}{
	{
		nome:          "Cliente VIP ganha 10% de desconto",
		clienteID:     "user_vip",
		valorOriginal: 23.00,
		ehVIP:         true,
		valorEsperado: 20.7,
	},
	{
		nome:          "Cliente comum não ganha desconto",
		clienteID:     "user_comum",
		valorOriginal: 45.00,
		ehVIP:         false,
		valorEsperado: 45.00,
	},
}

func (m *MockVip) EhhVip(clienteID string) bool {
	args := m.Called(clienteID)
	return args.Bool(0)
}

func TestCalcular(t *testing.T) {
	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			m := new(MockVip)
			m.On("EhhVip", tt.clienteID).Return(tt.ehVIP)

			c := CalculadoraDesconto{VipService: m}

			valor := c.Calcular(tt.clienteID, tt.valorOriginal)

			assert.Equal(t, tt.valorEsperado, valor)
			m.AssertExpectations(t)
		})
	}
}
