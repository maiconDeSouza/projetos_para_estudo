package frete

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCapital struct {
	mock.Mock
}

var testes = []struct {
	nome      string
	cidade    string
	ehCapital bool
	retorno   float64
}{
	{nome: "Eh Capital", cidade: "São Paulo", ehCapital: true, retorno: 10},
	{nome: "Não eh Capital", cidade: "Ferraz de Vaconcelos", ehCapital: false, retorno: 30},
}

func (m *MockCapital) EhCapital(cidade string) bool {
	args := m.Called(cidade)

	return args.Bool(0)
}

func TestXxx(t *testing.T) {
	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			mock := new(MockCapital)
			mock.On("EhCapital", tt.cidade).Return(tt.ehCapital)

			c := ServicoFrete{Calc: mock}

			f := c.CalcularFrete(tt.cidade)

			assert.Equal(t, tt.retorno, f)
			mock.AssertExpectations(t)
		})
	}
}
