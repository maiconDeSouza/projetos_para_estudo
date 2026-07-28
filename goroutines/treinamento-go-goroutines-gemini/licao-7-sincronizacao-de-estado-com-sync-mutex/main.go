package main

import (
	"fmt"
	"sync"
)

type Estoque struct {
	quantidade int
	mu         sync.RWMutex
}

func (e *Estoque) AddQuantidade(q int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if q <= 0 {
		fmt.Println("É impossível adicionar 0 ou menos itens no estoque")
		return
	}
	e.quantidade += q
}

func (e *Estoque) Comprar(q int, nome string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.quantidade < q {
		fmt.Printf("Venda recusada! Estoque insuficiente. [cliente %s]\n", nome)
		return
	}

	e.quantidade -= q

	fmt.Printf("Venda realizada! [cliente %s] Restam %d itens.\n", nome, e.quantidade)
}

func (e *Estoque) ConsultarEstoque() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.quantidade
}

var nomes = []string{
	"Ana", "Bruno", "Carlos", "Daniela", "Eduardo",
	"Fernanda", "Gabriel", "Helena", "Igor", "Juliana",
	"Lucas", "Mariana", "Nicolas", "Olivia", "Pedro",
}

func main() {
	estoque := Estoque{}
	estoque.AddQuantidade(100)
	wg := sync.WaitGroup{}

	for _, nome := range nomes {
		wg.Go(func() {
			estoque.Comprar(10, nome)
		})
	}

	wg.Wait()

}
