package pedidos

import (
	"context"
	"fmt"
	"projeto-processador-concorrente-de-pedidos-de-e-commerce/numeros"
	"sync"
	"time"
)

type Pedido struct {
	NomeCliente string
	Quantidade  uint
	Estoque     *Estoque
}

type PedidoResposta struct {
	Msg        string
	Err        error
	Quantidade uint
}

func (p *Pedido) ProcessarPagamentoEPedido(ctx context.Context, qtd uint, nome string) PedidoResposta {
	ch := make(chan PedidoResposta, 1)
	ps := PedidoResposta{}
	go func() {
		tempo := time.Duration(numeros.Aleatorios(10)) * time.Second
		time.Sleep(tempo)
		ps = p.Estoque.Vender(qtd, nome)

		ch <- ps
	}()

	select {
	case ps := <-ch:
		return ps
	case <-ctx.Done():
		ps.Err = ctx.Err()
		return ps
	}
}

type Estoque struct {
	mu         sync.RWMutex
	quantidade uint
}

func (e *Estoque) Vender(qdt uint, nome string) PedidoResposta {
	e.mu.Lock()
	defer e.mu.Unlock()

	ps := PedidoResposta{}

	if e.quantidade == 0 {
		ps.Msg = ""
		ps.Err = fmt.Errorf("[%s] Acabou o produto", nome)
		ps.Quantidade = e.quantidade
		return ps
	}

	if qdt > e.quantidade {
		ps.Msg = ""
		ps.Err = fmt.Errorf("Quantidade insufiente para a compra de [%s]", nome)
		ps.Quantidade = e.quantidade
		return ps
	}

	if qdt <= 0 {
		ps.Msg = ""
		ps.Err = fmt.Errorf("[%s] você não pode comprar zero itens ou menos", nome)
		ps.Quantidade = e.quantidade
		return ps
	}

	e.quantidade -= qdt

	ps.Msg = fmt.Sprintf("[%s] comprou %d itens", nome, qdt)
	ps.Err = nil
	ps.Quantidade = e.quantidade
	return ps

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
