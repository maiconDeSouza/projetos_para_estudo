package pedidos

import (
	"context"
	"fmt"
	"sync"
)

type Pedido struct {
	NomeCliente string
	Quantidade  uint
	Estoque     *Estoque
}

type PedidoResposta struct {
	msg        string
	err        error
	quantidade uint
}

func (p *Pedido) processarPagamentoEPedido(ctx context.Context)

type Estoque struct {
	mu         sync.RWMutex
	quantidade uint
}

func (e *Estoque) Vender(qdt uint, nome string) PedidoResposta {
	e.mu.Lock()
	defer e.mu.Unlock()

	ps := PedidoResposta{}

	if qdt > e.quantidade {
		ps.msg = ""
		ps.err = fmt.Errorf("Quantidade insufiente para a compra de [%s]", nome)
		ps.quantidade = e.quantidade
		return ps
	}

	if qdt <= 0 {
		ps.msg = ""
		ps.err = fmt.Errorf("[%s] você não pode comprar zero itens ou menos", nome)
		return ps
	}

	e.quantidade -= qdt

	return fmt.Sprintf("[%s] comprou %d itens", nome, qdt), nil
}

func (e *Estoque) Quantidade() uint {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.quantidade
}

func NewEstoque(qtd uint) *Estoque {
	return &Estoque{quantidade: qtd}
}

func NewPedido(nome string, qtd uint, estoque *Estoque) *Pedido {
	return &Pedido{
		NomeCliente: nome,
		Estoque:     estoque,
		Quantidade:  qtd,
	}
}
