package caixa

import (
	"errors"
	"fmt"
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

	fmt.Println(saldo, valor)

	if saldo < valor {
		return errors.New("Saldo insuficiente!")
	}

	saldo -= valor

	return nil
}
