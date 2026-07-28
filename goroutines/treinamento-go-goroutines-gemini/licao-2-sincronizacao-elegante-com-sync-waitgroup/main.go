package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"https://golang.org",
	"https://google.com",
	"https://github.com",
	"https://golangbr.org",
}

func verificarServidor(url string) {
	fmt.Printf("Verificando [%s]...\n", url)
	time.Sleep(time.Second)
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Servidor [%s] OFF!\n", url)
		return
	}

	fmt.Printf("Servidor [%s] OK!\n", url)
}

func main() {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Go(func() {
			verificarServidor(url)
		})
	}

	wg.Wait()

	fmt.Println("Checagem finalizada!")
}
