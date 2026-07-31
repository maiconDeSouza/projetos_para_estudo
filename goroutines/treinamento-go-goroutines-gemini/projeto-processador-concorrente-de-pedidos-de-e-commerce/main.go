package main

import (
	"context"
	"fmt"
	"projeto-processador-concorrente-de-pedidos-de-e-commerce/numeros"
	"projeto-processador-concorrente-de-pedidos-de-e-commerce/pedidos"
	"sync"
	"time"
)

func worker(pedidos <-chan pedidos.Pedido, resposta chan<- string, wk int, e *pedidos.Estoque) {
	for p := range pedidos {
		tempo := time.Duration(10) * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), tempo)

		ps := p.ProcessarPagamentoEPedido(p.ID, wk, ctx, p.Quantidade, p.NomeCliente)
		cancel()

		if ps.Err != nil {
			resposta <- ps.Err.Error()
			continue
		}

		n := numeros.Aleatorios(10)
		if n == 5 {
			qtd := numeros.Aleatorios(3)
			e.AddQuantidade(qtd)
		}

		resposta <- ps.Msg
	}
}

var nomes = []string{
	"Ana", "Bruno", "Carlos", "Daniela", "Eduardo",
	"Fernanda", "Gabriel", "Helena", "Igor", "Júlia",
	"Lucas", "Mariana", "Nicolas", "Olívia", "Pedro",
	"Rafaela", "Samuel", "Tatiane", "Vitor", "Yasmin",
}

func main() {
	wg := sync.WaitGroup{}
	totalDeCompras := numeros.Aleatorios(30)
	p := make(chan pedidos.Pedido, totalDeCompras)
	res := make(chan string, totalDeCompras)
	e := pedidos.NewEstoque(30)

	for i := 1; i <= totalDeCompras; i++ {
		index := numeros.Aleatorios(len(nomes)) - 1
		qtd := numeros.Aleatorios(10)
		cliente := nomes[index]
		novoPedido := pedidos.NewPedido(i, cliente, uint(qtd), e)
		p <- *novoPedido
	}

	close(p)

	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			worker(p, res, i, e)
		})
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	fmt.Printf("Quantidades de compras: %d para %d itens.\n", totalDeCompras, e.Quantidade())

	for r := range res {
		fmt.Println(r)
	}

	fmt.Printf("Total de Produtos em estoque %d", e.Quantidade())
}
