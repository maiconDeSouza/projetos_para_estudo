package caixa

import (
	"errors"
)

type BancoService interface {
	ConsultarSaldo(contaID string) (float64, error)
}

type CaixaEletronico struct {
	Banco BancoService
}

func (c *CaixaEletronico) Sacar(contaID string, valor float64) error {
	saldo, err := c.Banco.ConsultarSaldo(contaID)
	if err != nil {
		return err
	}

	if saldo < valor {
		return errors.New("Saldo insuficiente!")
	}

	saldo -= valor

	return nil
}
