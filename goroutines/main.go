package main

import (
	"fmt"
	"time"
)

func somaGigante() uint64 {
	var total uint64

	for i := uint64(1); i <= 1_000_000_000; i++ {
		total += i
	}

	return total
}

func somaIntervalo(inicio, fim uint64) uint64 {
	var total uint64

	for i := inicio; i <= fim; i++ {
		total += i
	}

	return total
}

// func main() {
// 	// inicio := time.Now()

// 	// resultado := somaGigante()

// 	// fmt.Println(resultado)
// 	// ch := make(chan uint64)
// 	// go func() {
// 	// 	ch <- somaIntervalo(1, 250_000_000)
// 	// }()

// 	// go func() {
// 	// 	ch <- somaIntervalo(250_000_001, 500_000_000)
// 	// }()

// 	// go func() {
// 	// 	ch <- somaIntervalo(500_000_001, 750_000_000)
// 	// }()

// 	// go func() {
// 	// 	ch <- somaIntervalo(750_000_001, 1_000_000_000)
// 	// }()
// 	// var total uint64

// 	// for i := 0; i < 4; i++ {
// 	// 	total += <-ch
// 	// }

// 	// fmt.Println(total)
// 	// fmt.Println(time.Since(inicio))
// 	fmt.Println(runtime.NumCPU())
// }

const limite uint64 = 50_000_000_000
const workers = 20

func main() {
	inicio := time.Now()

	ch := make(chan uint64)

	tamanhoBloco := limite / workers

	for i := uint64(0); i < workers; i++ {
		inicioBloco := i*tamanhoBloco + 1
		fimBloco := (i + 1) * tamanhoBloco

		if i == workers-1 {
			fimBloco = limite
		}

		go func(inicio, fim uint64) {
			ch <- somaIntervalo(inicio, fim)
		}(inicioBloco, fimBloco)
	}

	var total uint64

	for i := 0; i < workers; i++ {
		total += <-ch
	}

	fmt.Println(total)
	fmt.Println(time.Since(inicio))
}
