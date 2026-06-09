package com

import (
	"fmt"
	"strings"
	"time"
)

func intToRoman(number int) string {
	if number <= 0 || number > 3999 {
		return "Fora do intervalo permitido (1 - 3999)"
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}
	return roman.String()
}

func loopInt() int {
	var total int = 0
	for i := 0; i <= 100; i++ {
		fmt.Println("volta ->", i)
		time.Sleep(time.Duration(500) * time.Millisecond)
		total += i
	}
	return total
}

func loopRoman() int {
	var totalRoamano int = 0
	for i := 0; i <= 50; i++ {
		fmt.Printf("%d -> %s\n", i, intToRoman(i))
		time.Sleep(time.Duration(500) * time.Millisecond)
		totalRoamano += i
	}
	return totalRoamano
}

func Com() {
	inicio := time.Now()

	chInt := make(chan int)
	chRoman := make(chan int)

	go func() {
		chInt <- loopInt()
	}()

	go func() {
		chRoman <- loopRoman()
	}()

	total := <-chInt
	totalRomano := <-chRoman

	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Total: %s\n", intToRoman(totalRomano))
	decorrido := time.Since(inicio)
	fmt.Printf("O processamento demorou: %s\n", decorrido)
}
