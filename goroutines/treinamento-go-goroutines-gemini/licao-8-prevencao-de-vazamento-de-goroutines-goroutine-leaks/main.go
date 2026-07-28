// package main

// import (
// 	"fmt"
// 	"time"
// )

// func processarTarefa() string {
// 	ch := make(chan string)

// 	go func() {
// 		// Simula um processamento demorado de 3 segundos
// 		time.Sleep(3 * time.Second)
// 		ch <- "Tarefa Concluída com Sucesso!"
// 	}()

// 	timer := time.NewTimer(time.Duration(3) * time.Second)
// 	defer timer.Stop()
// 	select {
// 	case resultado := <-ch:
// 		return resultado
// 	case <-timer.C:
// 		return "Erro: Tempo limite excedido!"
// 	}
// }

// func main() {
// 	status := processarTarefa()
// 	fmt.Println("Status:", status)

//		// Espera um pouco para simular o programa continuando
//		time.Sleep(5 * time.Second)
//	}
package main

import (
	"fmt"
	"time"
)

func processarTarefa() string {
	ch := make(chan string, 1)

	go func() {
		// Simula um processamento demorado de 3 segundos
		time.Sleep(3 * time.Second)
		ch <- "Tarefa Concluída com Sucesso!"
	}()
	timer := time.NewTimer(time.Duration(1) * time.Second)
	defer timer.Stop()
	select {
	case resultado := <-ch:
		return resultado
	case <-timer.C:
		return "Erro: Tempo limite excedido!"
	}
}

func main() {
	status := processarTarefa()
	fmt.Println("Status:", status)

	// Espera um pouco para simular o programa continuando
	time.Sleep(4 * time.Second)
}
