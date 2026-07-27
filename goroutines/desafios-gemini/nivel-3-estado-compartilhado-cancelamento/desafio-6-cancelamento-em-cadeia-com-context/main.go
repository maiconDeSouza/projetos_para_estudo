package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Resultado struct {
	Servico string
	Dado    string
}

func buscarServico(ctx context.Context, nome string, ch chan<- Resultado) {
	delay := rand.Intn(3) + 1

	select {
	case <-time.After(time.Second * time.Duration(delay)):
		ch <- Resultado{Servico: nome, Dado: "Dados de " + nome}
	case <-ctx.Done():
		fmt.Printf("[%s] Operação cancelada! Liberando recursos...\n", nome)
		return
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan Resultado, 3)

	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			buscarServico(ctx, fmt.Sprintf("Servico %d", i), ch)
		})
	}

	primeiroResultado := <-ch
	fmt.Printf("\n🏆 Ganhou: %s -> %s\n\n", primeiroResultado.Servico, primeiroResultado.Dado)

	cancel()

	wg.Wait()
	time.Sleep(time.Duration(5) * time.Second)
}
