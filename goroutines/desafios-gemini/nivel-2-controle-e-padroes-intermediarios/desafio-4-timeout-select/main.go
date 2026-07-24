package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	resp := make(chan string, 10)

	wg.Go(func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			resp <- "Sucesso!"
		}
	})

	go func() {
		wg.Wait()
		close(resp)
	}()

	for {
		select {
		case msg, ok := <-resp:
			if !ok {
				return
			}
			fmt.Printf("Resposta recebida: %s\n", msg)
		case <-time.After(time.Duration(2) * time.Second):
			fmt.Println("Erro: Operação expirou (Timeout)!")
		}
	}

}
