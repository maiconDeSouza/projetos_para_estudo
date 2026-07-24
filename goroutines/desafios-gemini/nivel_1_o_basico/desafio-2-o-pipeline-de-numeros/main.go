package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Go(func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	})

	go func() {
		for c := range ch {
			fmt.Println(c * 2)
		}
	}()

	wg.Wait()
	close(ch)

}
