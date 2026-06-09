package main

import "fmt"

func filter[T any](s []T, f func(i int, value T, s []T) bool) []T {
	var newSlice []T
	for i, v := range s {
		if f(i, v, s) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func mapGo[T any](s []T, f func(i int, value T, s []T) T) []T {
	var newSlice []T

	for i, v := range s {
		newV := f(i, v, s)
		newSlice = append(newSlice, newV)
	}
	return newSlice
}

func reduce[T any](s []T, initialValue T, f func(acc T, value T, s []T) T) T {
	acc := initialValue

	for _, v := range s {
		acc = f(acc, v, s)
	}
	return acc
}

func testeFilter(i int, value int, s []int) bool {
	return value%2 == 0
}

func testeMap(i int, value int, s []int) int {
	return value * 2
}

func testeReduce(acc int, value int, s []int) int {
	return acc + value
}

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	f := filter(numbers, testeFilter)
	fmt.Println("Pares filtrados:", f)

	m := mapGo(numbers, testeMap)
	fmt.Println("Slice dobrado:", m)

	r := reduce(numbers, 0, testeReduce)
	fmt.Println("Reduce:", r)

}
