package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var gate sync.RWMutex
var wg sync.WaitGroup

type Contador struct {
	valor uint
}

func (c *Contador) Incrementando() {
	gate.Lock()
	defer gate.Unlock()
	c.valor++
}

func (c Contador) PegarValor() int {
	return int(c.valor)
}

func main() {
	contador := Contador{valor: 0}

	for i := 1; i <= 1000; i++ {
		wg.Go(func() {
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			contador.Incrementando()
		})
	}

	wg.Wait()

	fmt.Printf("O valor atual do contador é %d", contador.PegarValor())

}
