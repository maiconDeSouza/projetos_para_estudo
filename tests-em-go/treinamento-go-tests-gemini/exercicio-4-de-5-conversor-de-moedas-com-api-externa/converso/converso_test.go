package converso

import (
	"errors"
	"testing"
)

var taxaBRLtoUSD float64 = 0.1932367
var taxaUSDtoBRL float64 = 5.0842000

type MockProvedorDeCotacao struct {
	chamada      bool
	err          error
	taxaBRLtoUSD float64
	taxaUSDtoBRL float64
}

func (m *MockProvedorDeCotacao) ObterTaxa(daMoeda, paraMoeda Moeda) (float64, error) {
	if m.err != ErrCotacao {
		return 0.00, ErrCotacao
	}

	if daMoeda == BRL {
		m.chamada = true
		return taxaBRLtoUSD, nil
	}

	if daMoeda == USD {
		m.chamada = true
		return taxaUSDtoBRL, nil
	}

	return 0.00, ErrCotacao
}

func TestConverter(t *testing.T) {
	t.Run("Valor menor ou igual a zero 🚫", func(t *testing.T) {
		mk := MockProvedorDeCotacao{chamada: false, err: nil, taxaBRLtoUSD: taxaBRLtoUSD, taxaUSDtoBRL: taxaUSDtoBRL}
		con := ConversorMoedas{Provedor: &mk}
		valor := 0.00

		valorCovertido, err := con.Converter(valor, BRL, USD)
		if valorCovertido != valor {
			t.Errorf("O valor esperado era [%.2f] e recebemos [%.2f]", valor, valorCovertido)
		}
		if !errors.Is(err, ErrValorInvalido) {
			t.Errorf("O valor esperado era [%v] e recebemos [%v]", ErrValorInvalido, err)
		}
	})

	t.Run("Mesma moeda origem/destino (garantindo que a API **não** foi chamada) 🔄", func(t *testing.T) {
		mk := MockProvedorDeCotacao{chamada: false, err: nil, taxaBRLtoUSD: taxaBRLtoUSD, taxaUSDtoBRL: taxaUSDtoBRL}
		con := ConversorMoedas{Provedor: &mk}
		valor := 23.00

		valorCovertido, err := con.Converter(valor, BRL, BRL)
		if valorCovertido != valor {
			t.Errorf("O valor esperado era [%.2f] e recebemos [%.2f]", valor, valorCovertido)
		}

		if err != nil {
			t.Errorf("O valor esperado era [%v] e recebemos [%v]", nil, err)
		}

		if mk.chamada {
			t.Error("Algo aconteceu de errado e função ObterTaxa foi chamada")
		}
	})

	t.Run("Falha na chamada da API de cotação ⚠️", func(t *testing.T) {
		mk := MockProvedorDeCotacao{chamada: true, err: ErrCotacao, taxaBRLtoUSD: taxaBRLtoUSD, taxaUSDtoBRL: taxaUSDtoBRL}
		con := ConversorMoedas{Provedor: &mk}
		valor := 23.00

		valorCovertido, err := con.Converter(valor, BRL, USD)
		if valorCovertido != valor {
			t.Errorf("O valor esperado era [%.2f] e recebemos [%.2f]", valor, valorCovertido)
		}
		if !errors.Is(err, ErrCotacao) {
			t.Errorf("O valor esperado era [%v] e recebemos [%v]", ErrCotacao, err)
		}
	})

	t.Run("Conversão realizada com sucesso 🟢", func(t *testing.T) {
		mk := MockProvedorDeCotacao{chamada: true, err: nil, taxaBRLtoUSD: taxaBRLtoUSD, taxaUSDtoBRL: taxaUSDtoBRL}
		con := ConversorMoedas{Provedor: &mk}
		valor := 23.00
		valorEsperado := valor * taxaBRLtoUSD

		valorCovertido, err := con.Converter(valor, BRL, USD)
		if valorCovertido != valorEsperado {
			t.Errorf("O valor esperado era [%.2f] e recebemos [%.2f]", valorEsperado, valorCovertido)
		}
		if err != nil {
			t.Errorf("O valor esperado era [%v] e recebemos [%v]", nil, err)
		}
	})
}
