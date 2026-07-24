package main

import (
	"fmt"
	"sync"
)

func helloWorld(wg *sync.WaitGroup) {
	for i := range 4 {
		wg.Go(func() {
			fmt.Printf("Goroutine %d executando\n", i+1)
		})
	}
}

func main() {
	var wg sync.WaitGroup
	helloWorld(&wg)

	wg.Wait()
}
