package main

import (
	"context"
	"fmt"
	"projeto-processador-concorrente-de-pedidos-de-e-commerce/numeros"
	"projeto-processador-concorrente-de-pedidos-de-e-commerce/pedidos"
	"sync"
	"time"
)

func worker(pedidos <-chan pedidos.Pedido, resposta chan<- string) {
	for p := range pedidos {
		tempo := time.Duration(10) * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), tempo)

		ps := p.ProcessarPagamentoEPedido(ctx, p.Quantidade, p.NomeCliente)
		cancel()

		if ps.Err != nil {
			resposta <- ps.Err.Error()
			return
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
	totalDeCompras := numeros.Aleatorios(100)
	p := make(chan pedidos.Pedido, totalDeCompras)
	res := make(chan string, totalDeCompras)
	e := pedidos.NewEstoque(30)

	for i := 1; i <= totalDeCompras; i++ {
		index := numeros.Aleatorios(len(nomes)) - 1
		qtd := numeros.Aleatorios(10)
		cliente := nomes[index]
		novoPedido := pedidos.NewPedido(cliente, uint(qtd), e)
		p <- *novoPedido
	}

	close(p)

	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			worker(p, res)
		})
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		fmt.Println(r)
	}

	fmt.Printf("Total de Produtos em estoque %d", e.Quantidade())
}
