package estoque

import (
	"errors"
	"testing"
)

type MockEstoque struct {
	err           error
	foiNotificado bool
}

func (m *MockEstoque) NotificarEstoqueBaixo(produtoID string, quantidadeAtual int) error {
	m.foiNotificado = true
	return m.err
}

func TestBaixarEstoque(t *testing.T) {
	t.Run("Quantidade comprada inválida (ex: `0` ou `-1`) 🚫", func(t *testing.T) {
		mock := MockEstoque{err: nil, foiNotificado: false}
		g := GerenciadorEstoque{Notificador: &mock}

		_, err := g.BaixarEstoque("a1", 23, -1)
		if mock.foiNotificado {
			t.Error("Não era para ter notificado o pessoal de compras")
		}
		if !errors.Is(err, ErrQTDInvalida) {
			t.Errorf("Era esperado o erro [%v] e recebemos o erro [%v]", ErrQTDInvalida, err)
		}
	})

	t.Run("Estoque insuficiente (comprar mais do que tem) ❌", func(t *testing.T) {
		mock := MockEstoque{err: nil, foiNotificado: false}
		g := GerenciadorEstoque{Notificador: &mock}

		_, err := g.BaixarEstoque("a1", 3, 5)
		if mock.foiNotificado {
			t.Error("Não era para ter notificado o pessoal de compras")
		}
		if !errors.Is(err, ErrEstoqueInsuficiente) {
			t.Errorf("Era esperado o erro [%v] e recebemos o erro [%v]", ErrEstoqueInsuficiente, err)
		}
	})

	t.Run("Baixa de estoque com sucesso e estoque alto (não deve notificar) 🟢", func(t *testing.T) {
		mock := MockEstoque{err: nil, foiNotificado: false}
		g := GerenciadorEstoque{Notificador: &mock}

		_, err := g.BaixarEstoque("a1", 23, 5)
		if mock.foiNotificado {
			t.Error("Não era para ter notificado o pessoal de compras")
		}
		if err != nil {
			t.Errorf("Não era esperado nenhum erro, mas recebemos [%v]", err)
		}
	})

	t.Run("Baixa de estoque com sucesso e estoque alto (não deve notificar) (testando quantidade) 🟢", func(t *testing.T) {
		mock := MockEstoque{err: nil, foiNotificado: false}
		g := GerenciadorEstoque{Notificador: &mock}

		qtd, err := g.BaixarEstoque("a1", 23, 5)
		if mock.foiNotificado {
			t.Error("Não era para ter notificado o pessoal de compras")
		}
		if err != nil {
			t.Errorf("Não era esperado nenhum erro, mas recebemos [%v]", err)
		}

		if qtd != 18 {
			t.Errorf("Quantidade errada, recebemos [%d], mas era esperedado 18", qtd)
		}
	})

	t.Run("Baixa de estoque com sucesso e estoque baixo (deve notificar) (sem erro de notificação) ⚠️", func(t *testing.T) {
		mock := MockEstoque{err: nil}
		g := GerenciadorEstoque{Notificador: &mock}

		_, err := g.BaixarEstoque("a1", 9, 5)
		if !mock.foiNotificado {
			t.Error("Foi chamada a função de notificação")
		}
		if err != nil {
			t.Errorf("Não era esperado nenhum erro, mas recebemos [%v]", err)
		}
	})

	t.Run("Baixa de estoque com sucesso e estoque baixo (deve notificar) (com erro de notificação) ⚠️", func(t *testing.T) {
		mock := MockEstoque{err: ErrFalhaMSG}
		g := GerenciadorEstoque{Notificador: &mock}

		_, err := g.BaixarEstoque("a1", 9, 5)
		if !mock.foiNotificado {
			t.Error("Foi chamada a função de notificação")
		}
		if !errors.Is(err, ErrFalhaMSG) {
			t.Errorf("Era esperado o erro [%v] e recebemos o erro [%v]", ErrFalhaMSG, err)
		}
	})
}
