package main

import (
	"fmt"
	"time"
)

func calcularSalarioComBonus(salarioBase float64, bonus float64, ch chan float64) {
	time.Sleep(time.Duration(1) * time.Second)

	ch <- salarioBase + bonus
}

func main() {
	ch := make(chan float64)

	go calcularSalarioComBonus(3000, 500, ch)

	msg := fmt.Sprintf("Salário total calculado: R$ %.2f", <-ch)

	fmt.Println(msg)
}
