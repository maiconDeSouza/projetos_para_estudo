package main

import (
	"fmt"
	"sync"
	"time"
)

func generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- i
		}
	}()

	return out
}

func worker(id int, in <-chan int) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for n := range in {
			r := n * 2
			time.Sleep(time.Duration(1) * time.Second)
			out <- fmt.Sprintf("Trabalhador %d entregou o número %d", id, r)
		}
	}()
	return out
}

func fanIn(channels ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	for _, ch := range channels {
		chnew := ch
		wg.Go(func() {
			for c := range chnew {
				out <- c
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := generate(100)

	wk1 := worker(1, in)
	wk2 := worker(2, in)
	wk3 := worker(3, in)
	wk4 := worker(4, in)

	finalChan := fanIn(wk1, wk2, wk3, wk4)

	for res := range finalChan {
		fmt.Println(res)
	}

}
