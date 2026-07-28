package main

import (
	"fmt"
	"time"
)

func processarPedido(id int) {
	fmt.Printf("Iniciando pedido %d\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Pedido %d concluído!\n", id)
}

func main() {
	go processarPedido(1)
	go processarPedido(2)
	go processarPedido(3)

	time.Sleep(time.Duration(5) * time.Second)
}
