package pedidos

import (
	"context"
	"fmt"
	"projeto-processador-concorrente-de-pedidos-de-e-commerce/numeros"
	"sync"
	"time"
)

type Pedido struct {
	ID          int
	NomeCliente string
	Quantidade  uint
	Estoque     *Estoque
}

type PedidoResposta struct {
	Msg        string
	Err        error
	Quantidade uint
}

func (p *Pedido) ProcessarPagamentoEPedido(id, wk int, ctx context.Context, qtd uint, nome string) PedidoResposta {
	ch := make(chan PedidoResposta, 1)
	ps := PedidoResposta{}
	go func() {
		tempo := time.Duration(numeros.Aleatorios(10)) * time.Second
		time.Sleep(tempo)
		res := p.Estoque.Vender(id, wk, qtd, nome)

		ch <- res
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

func (e *Estoque) Vender(id, wk int, qdt uint, nome string) PedidoResposta {
	e.mu.Lock()
	defer e.mu.Unlock()

	ps := PedidoResposta{}

	if e.quantidade == 0 {
		ps.Msg = ""
		ps.Err = fmt.Errorf("[%d] --- Trabalhador [%d] --- [%s] tentou comprar %d, mas o estoque acabou - estoque:[%d]", id, wk, nome, qdt, e.quantidade)
		ps.Quantidade = e.quantidade
		return ps
	}

	if qdt > e.quantidade {
		ps.Msg = ""
		ps.Err = fmt.Errorf("[%d] --- Trabalhador [%d] --- [%s]tentou comprar %d, mas só tem em estoque:[%d]", id, wk, nome, qdt, e.quantidade)
		ps.Quantidade = e.quantidade
		return ps
	}

	if qdt <= 0 {
		ps.Msg = ""
		ps.Err = fmt.Errorf("[%d] --- Trabalhador [%d] --- [%s] você não pode comprar zero itens ou menos", id, wk, nome)
		ps.Quantidade = e.quantidade
		return ps
	}

	e.quantidade -= qdt

	ps.Msg = fmt.Sprintf("[%d] --- Trabalhador [%d] --- [%s] comprou %d itens - estoque:[%d]", id, wk, nome, qdt, e.quantidade)
	ps.Err = nil
	ps.Quantidade = e.quantidade
	return ps

}

func (e *Estoque) Quantidade() uint {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.quantidade
}

func (e *Estoque) AddQuantidade(n int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.quantidade += uint(n)
}

func NewEstoque(qtd uint) *Estoque {
	return &Estoque{quantidade: qtd}
}

func NewPedido(id int, nome string, qtd uint, estoque *Estoque) *Pedido {
	return &Pedido{
		ID:          id,
		NomeCliente: nome,
		Estoque:     estoque,
		Quantidade:  qtd,
	}
}
