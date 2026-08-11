package estoque

import "errors"

type Notificador interface {
	NotificarEstoqueBaixo(produtoID string, quantidadeAtual int) error
}

type GerenciadorEstoque struct {
	Notificador Notificador
}

var ErrQTDInvalida = errors.New("quantidade invalida")
var ErrEstoqueInsuficiente = errors.New("estoque insuficiente")
var ErrFalhaMSG = errors.New("falha ao notificar estoque baixo")

func (g *GerenciadorEstoque) BaixarEstoque(produtoID string, quantidadeAtual int, quantidadeComprada int) (int, error) {
	if quantidadeComprada <= 0 {
		return quantidadeAtual, ErrQTDInvalida
	}

	if quantidadeAtual < quantidadeComprada {
		return quantidadeAtual, ErrEstoqueInsuficiente
	}

	quantidadeAtual -= quantidadeComprada

	if quantidadeAtual < 5 {
		err := g.Notificador.NotificarEstoqueBaixo(produtoID, quantidadeAtual)
		if err != nil {
			return quantidadeAtual, err
		}
		return quantidadeAtual, nil
	}

	return quantidadeAtual, nil
}
