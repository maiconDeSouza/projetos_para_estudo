package main

import (
	"fmt"
	"sync"
	"time"
)

func display(l string) {
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(l)
}

func main() {
	wg := sync.WaitGroup{}
	alpha := "abcdefghijklmnopqrstuvwxyz"

	for _, l := range alpha {
		wg.Go(func() {
			display(string(l))
		})

	}

	wg.Wait()
}
