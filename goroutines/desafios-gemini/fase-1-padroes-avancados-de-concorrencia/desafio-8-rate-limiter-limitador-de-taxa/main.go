package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	const numeroRequestes = 20
	const limite = time.Second / time.Duration(5)

	requestes := make(chan int)

	for i := 1; i <= numeroRequestes; i++ {
		wg.Go(func() {
			requestes <- i
		})
	}

	go func() {
		wg.Wait()
		close(requestes)
	}()

	ticker := time.NewTicker(limite)
	defer ticker.Stop()

	for req := range requestes {
		<-ticker.C
		fmt.Printf("[%s] Processando requisição %d\n", time.Now().Format("15:04:05.000"), req)
	}
}
