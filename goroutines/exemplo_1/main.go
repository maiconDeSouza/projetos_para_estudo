package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(letter string, channel chan string) {
	channel <- string(letter)
	time.Sleep(time.Duration(2) * time.Second)
}

func main() {
	var wg sync.WaitGroup
	var channel = make(chan string)
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	startTime := time.Now()

	for _, s := range alpha {
		wg.Go(func() {
			escrever(string(s), channel)
		})

	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	for letter := range channel {
		fmt.Println(letter)
	}

	t := time.Since(startTime)
	fmt.Printf("A aplicação levou %.2f", t.Seconds())
}
