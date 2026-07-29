package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, faturas <-chan int, resultados chan<- string) {
	for faturaID := range faturas {
		time.Sleep(time.Duration(1) * time.Second)
		resultados <- fmt.Sprintf("Fatura [%d] processada pelo Worker [%d]", faturaID, id)
	}
}

var faturasID = []int{1, 2, 3, 4, 5}

func main() {
	faturas := make(chan int, len(faturasID))
	resultados := make(chan string, len(faturasID))

	for _, f := range faturasID {
		faturas <- f
	}
	close(faturas)

	wg := sync.WaitGroup{}

	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			worker(i, faturas, resultados)
		})
	}

	go func() {
		wg.Wait()
		close(resultados)
	}()

	for r := range resultados {
		fmt.Println(r)
	}
}
