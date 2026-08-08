package caixa

import (
	"errors"
	"testing"
)

type MockBancoService struct {
	contaID bool
	saldo   float64
}

func (m *MockBancoService) ConsultarSaldo(contaID string) (float64, error) {
	if !m.contaID {
		return 0.00, errors.New("conta inexistente!")
	}
	return m.saldo, nil
}

func TestSacar(t *testing.T) {
	t.Run("Sucesso!", func(t *testing.T) {
		mock := MockBancoService{contaID: true, saldo: 1000}
		caixa := CaixaEletronico{Banco: &mock}
		saque := 500.00
		err := caixa.Sacar("a2", saque)

		if err != nil {
			t.Errorf("Não era esperado um erro quando a conta é %v quando o saldo é de %.2f e saque era de %.2f", mock.contaID, mock.saldo, saque)
		}
	})

	t.Run("Conta Inválida", func(t *testing.T) {
		mock := MockBancoService{contaID: false, saldo: 1000}
		caixa := CaixaEletronico{Banco: &mock}
		saque := 500.00

		err := caixa.Sacar("a2", saque)

		if err == nil {
			t.Errorf("Era esperado um erro quando a conta é %v quando o saldo é de %.2f e saque era de %.2f", mock.contaID, mock.saldo, saque)
		}
	})

	t.Run("Saldo insufiente!", func(t *testing.T) {
		mock := MockBancoService{contaID: false, saldo: 400}
		caixa := CaixaEletronico{Banco: &mock}
		saque := 500.00

		err := caixa.Sacar("a2", saque)

		if err == nil {
			t.Errorf("Era esperado um erro de saldo insufienciente para saldo %.2f e um saque de %.2f", mock.saldo, saque)
		}
	})
}
