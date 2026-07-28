package main

import (
	"fmt"
	"sync"
	"time"
)

func cotarDolar(chDolar chan string) {
	time.Sleep(time.Duration(1) * time.Second)
	chDolar <- "USD: R$ 5.00"
}

func cotarEuro(chEuro chan string) {
	time.Sleep(time.Duration(2) * time.Second)
	chEuro <- "EUR: R$ 5.50"
}

func main() {
	var wg sync.WaitGroup

	chDolar := make(chan string)
	chEuro := make(chan string)

	wg.Go(func() {
		cotarDolar(chDolar)
	})

	wg.Go(func() {
		cotarEuro(chEuro)
	})

	timer := time.NewTimer(time.Duration(3) * time.Second)
	defer timer.Stop()

	for i := 1; i <= 2; i++ {
		timer.Reset(time.Duration(3) * time.Second)
		select {
		case msg := <-chDolar:
			fmt.Println(msg)
		case msg := <-chEuro:
			fmt.Println(msg)
		case <-timer.C:
			fmt.Println("Nenhuma cotação respondeu!")
			return
		}
	}
	wg.Wait()
}
