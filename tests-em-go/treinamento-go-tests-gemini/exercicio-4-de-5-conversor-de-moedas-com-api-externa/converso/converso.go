package converso

import "errors"

type Moeda string

const (
	USD Moeda = "USD"
	BRL Moeda = "BRL"
)

var ErrValorInvalido = errors.New("valor invalido")
var ErrCotacao = errors.New("Falha na Cotação")

type ProvedorCotacao interface {
	ObterTaxa(daMoeda, paraMoeda Moeda) (float64, error)
}

type ConversorMoedas struct {
	Provedor ProvedorCotacao
}

func (c *ConversorMoedas) Converter(valor float64, daMoeda, paraMoeda Moeda) (float64, error) {
	if valor <= 0 {
		return valor, ErrValorInvalido
	}

	if daMoeda == paraMoeda {
		return valor, nil
	}

	taxa, err := c.Provedor.ObterTaxa(daMoeda, paraMoeda)
	if err != nil {
		return valor, err
	}

	valorCovertido := valor * taxa
	return valorCovertido, nil
}
